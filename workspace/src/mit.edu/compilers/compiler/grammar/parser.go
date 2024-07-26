package grammar

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
)

func isTtypeVisible(ttype int) bool {
	return ttype == IDENTIFIER
}

func newLexer(fin io.Reader, onErr lexmachine.Action) *golex {
	lexer := lexmachine.NewLexer()

	tokmap := make(map[string]int)
	for id, name := range yyToknames {
		tokmap[name] = id
	}
	emitTokenAction := func(ttypeName string) lexmachine.Action {
		return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
			return s.Token(tokmap[ttypeName], string(m.Bytes), m), nil
		}
	}

	skip := func(s *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
		return nil, nil
	}

	lexer.Add([]byte("class"), emitTokenAction("CLASS"))

	lexer.Add([]byte("'(\\\\[n\"]|[^\"])'"), emitTokenAction("CHARLITERAL"))
	lexer.Add([]byte("\"(\\\\[n\"]|[^\"])*\""), emitTokenAction("STRINGLITERAL"))

	lexer.Add([]byte("[a-zA-Z]+"), emitTokenAction("IDENTIFIER"))

	lexer.Add([]byte("{"), emitTokenAction("LCURLY"))
	lexer.Add([]byte("}"), emitTokenAction("RCURLY"))

	// single-line comment
	lexer.Add([]byte("//[^\\n]*\\n"), skip)
	// whitespace
	lexer.Add([]byte("[ \\n]"), skip)

	lexer.Add([]byte("."), onErr)

	err := lexer.CompileNFA()
	if err != nil {
		panic(err)
	}
	code, err := ioutil.ReadAll(fin)
	if err != nil {
		panic(err)
	}
	scanner, err := lexer.Scanner(code)
	if err != nil {
		panic(err)
	}
	return &golex{Scanner: scanner}
}

func newPrintLexErrorAction(fout io.Writer, name string, numErrors *int) lexmachine.Action {
	return func(s *lexmachine.Scanner, match *machines.Match) (interface{}, error) {
		fmt.Fprintf(
			fout,
			"%s line %d:%d: unexpected char: '%s'\n",
			name,
			match.StartLine,
			match.StartColumn,
			match.Bytes,
		)
		*numErrors++
		return nil, nil
	}
}

func Lex(fin io.Reader, fout io.Writer, name string) {
	numErrors := 0
	onErr := newPrintLexErrorAction(fout, name, &numErrors)
	lexer := newLexer(fin, onErr)
	for tok, err, eos := lexer.Next(); !eos; tok, err, eos = lexer.Next() {
		if err != nil {
			panic(err)
		}
		token := tok.(*lexmachine.Token)
		ttype := token.Type
		if isTtypeVisible(toYaccTtype(ttype)) {
			ttypeStr := yyToknames[ttype]
			fmt.Fprintf(
				fout,
				"%d %s %s\n",
				token.StartLine,
				ttypeStr,
				token.Lexeme,
			)
		} else {
			fmt.Fprintf(
				fout,
				"%d %s\n",
				token.StartLine,
				token.Lexeme,
			)
		}
	}
	if numErrors > 0 {
		os.Exit(1)
	}
}

func Parse(fin io.Reader, fout io.Writer, name string, debug bool) *Node {
	numErrors := 0
	onErr := newPrintLexErrorAction(fout, name, &numErrors)
	lexer := newLexer(fin, onErr)
	yyErrorVerbose = true
	if debug {
		yyDebug = 2
	}
	yyParse(lexer)
	if numErrors > 0 {
		os.Exit(1)
	}
	return ProgramAst
}

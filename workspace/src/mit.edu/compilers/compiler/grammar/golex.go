package grammar

import (
	"fmt"

	"github.com/timtadh/lexmachine"
)

type golex struct {
	*lexmachine.Scanner
}

func toYaccTtype(ttype int) int {
	return ttype + yyPrivate - 1
}

func (g *golex) Lex(lval *yySymType) (tokenType int) {
	s := g.Scanner
	tok, err, eof := s.Next()
	if err != nil {
		g.Error(err.Error())
	} else if eof {
		return -1 // signals EOF to goyacc's yyParse
	}
	lval.Token = tok.(*lexmachine.Token)
	return toYaccTtype(lval.Token.Type)
}

func (g *golex) Error(message string) {
	panic(fmt.Errorf(message))
}

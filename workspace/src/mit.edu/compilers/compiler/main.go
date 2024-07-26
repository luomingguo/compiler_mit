package main

import (
	"os"
	"path/filepath"
	"strings"

	"mit.edu/compilers/compiler/grammar"
)

func main() {
	var err error

	cli := NewCLI()
	cli.ParseArgs(os.Args[1:], []string{})

	fout := os.Stdout
	if cli.outfile != "" {
		fout, err = os.Create(cli.outfile)
		if err != nil {
			panic(err)
		}
		defer fout.Close()
	}

	infileShortName := "stdin"
	fin := os.Stdin
	if cli.infile != "" {
		fin, err = os.Open(cli.infile)
		defer fin.Close()
		if err != nil {
			panic(err)
		}
		infileShortName = strings.Split(filepath.Base(cli.infile), ".")[0]
	}

	if cli.target == CLI_TARGET_SCAN {
		grammar.Lex(fin, fout, infileShortName)
		return
	}

	grammar.Parse(fin, fout, infileShortName, cli.debug)
}

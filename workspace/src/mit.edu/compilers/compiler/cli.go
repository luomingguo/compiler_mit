package main

import (
	"strings"
	"log"
)

const (
	CLI_TARGET_DEFAULT = iota
	CLI_TARGET_SCAN
	CLI_TARGET_PARSE
	CLI_TARGET_INTER
	CLI_TARGET_ASSEMBLY
)

type CLI struct {
	infile string
	outfile string
	target int
	debug bool
	opts []bool
}

func NewCLI() *CLI {
	return new(CLI)
}

func (cli *CLI) ParseArgs(args []string, optnames []string) {
	var extras []string
	targetStr := ""
	optsListStr := ""

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "--target=") {
			targetStr = arg[9:]
		} else if arg == "-t" {
			if i + 1 >= len(args) {
				log.Fatal("no target specified with -t")
			}
			targetStr = args[i + 1]
			i++
		} else if strings.HasPrefix(arg, "--outfile=") {
			cli.outfile = arg[10:]
		} else if arg == "-o" {
			if i + 1 >= len(args) {
				log.Fatal("no target specified with -o")
			}
			cli.outfile = args[i + 1]
			i++
		} else if strings.HasPrefix(arg, "--opt=") {
			optsListStr = arg[6:]
		} else if arg == "-O" {
			if i + 1 >= len(args) {
				log.Fatal("no target specified with -O")
			}
			optsListStr = args[i + 1]
			i++
		} else if arg == "--debug" || arg == "-d" {
			cli.debug = true
		} else {
			extras = append(extras, arg)
		}
	}

	targetStr = strings.ToLower(targetStr)
	switch targetStr {
		case "scan": cli.target = CLI_TARGET_SCAN
		case "parse": cli.target = CLI_TARGET_PARSE
		case "inter": cli.target = CLI_TARGET_INTER
		case "assembly": cli.target = CLI_TARGET_ASSEMBLY
		default: cli.target = CLI_TARGET_DEFAULT
	}

	cli.opts = make([]bool, len(optnames))
	optsArgs := strings.Split(optsListStr, ",")
	for _, optsArg := range optsArgs {
		if optsArg == "all" {
			for i := 0; i < len(optnames); i++ {
				cli.opts[i] = true
			}
			continue
		}
		for i := 0; i < len(optnames); i++ {
			if optsArg == optnames[i] {
				cli.opts[i] = true
				break
			} else if optsArg == "-" + optnames[i] {
				cli.opts[i] = false
				break
			}
		}
	}

	if len(extras) > 0 {
		cli.infile = extras[0]
	}
}

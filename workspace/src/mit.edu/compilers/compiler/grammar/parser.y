%{
package grammar

import (
	"github.com/timtadh/lexmachine"
)

var ProgramAst *Node
%}

%union {
	Token *lexmachine.Token
	Node *Node
}

%token<Token> CLASS IDENTIFIER
%token<Token> CHARLITERAL STRINGLITERAL
%token<Token> LCURLY RCURLY

%type<Node> prog block

%%

prog: CLASS IDENTIFIER block { ProgramAst = NewNode(NTYPE_NODE).AddChild(NewTokenNode($2)).AddChild($3) }

block: LCURLY RCURLY { $$ = NewNode(NTYPE_NODE) }

%%

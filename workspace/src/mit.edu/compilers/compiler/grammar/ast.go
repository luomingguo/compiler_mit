package grammar

import (
	"github.com/timtadh/lexmachine"
)

const (
	NTYPE_TOKEN = iota
	NTYPE_NODE
)

type Node struct {
	Type int
	Token *lexmachine.Token
	Children []*Node
}

func NewTokenNode(t *lexmachine.Token) *Node {
	return &Node{
		Type: NTYPE_TOKEN,
		Token: t,
	}
}

func NewNode(t int) *Node {
	return &Node{Type: t}
}

func (n *Node) AddChild(on *Node) *Node {
	return &Node{
		Type: n.Type,
		Children: append(n.Children, on),
	}
}

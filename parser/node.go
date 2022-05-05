package parser

import (
	"github.com/Div9851/mini-go/tokenizer"
)

type NodeKind int

const (
	ND_NUM NodeKind = iota
	ND_VAR
	ND_ADD
	ND_SUB
	ND_MUL
	ND_DIV
	ND_MOD
	ND_UNARY_MINUS
	ND_EXPR_LIST
	ND_EMPTY_STMT
	ND_EXPR_STMT
	ND_ASSIGN_STMT
)

// AST node
type Node struct {
	Kind NodeKind
	Tok  *tokenizer.Token // representative token

	Lhs *Node // Left-hand side
	Rhs *Node // Right-hand side

	Num     int     // Used if Kind == ND_NUM
	VarName string  // Used if kind == ND_VAR
	Offset  int     // Used if Kind == ND_VAR
	Exprs   []*Node // Used if Kind == ND_EXPR_LIST
}

func newNode(kind NodeKind, tok *tokenizer.Token) *Node {
	node := new(Node)
	node.Kind = kind
	node.Tok = tok
	return node
}

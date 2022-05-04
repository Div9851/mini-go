package parser

import (
	"bytes"

	"github.com/Div9851/mini-go/tokenizer"
)

func Parse(tokens []*tokenizer.Token) *Node {
	node, _ := expr(tokens)
	return node
}

func expr(tokens []*tokenizer.Token) (*Node, []*tokenizer.Token) {
	return addExpr(tokens)
}

func addExpr(tokens []*tokenizer.Token) (*Node, []*tokenizer.Token) {
	lhs, rest := mulExpr(tokens)
	tokens = rest
	for {
		var node *Node
		if isPunct(tokens, "+") {
			node = newNode(ND_ADD)
		} else if isPunct(tokens, "-") {
			node = newNode(ND_SUB)
		} else {
			break
		}
		tokens = tokens[1:]
		node.Lhs = lhs
		rhs, rest := mulExpr(tokens)
		node.Rhs = rhs
		tokens = rest
		lhs = node
	}
	return lhs, tokens
}

func mulExpr(tokens []*tokenizer.Token) (*Node, []*tokenizer.Token) {
	lhs, rest := unaryExpr(tokens)
	tokens = rest
	for {
		var node *Node
		if isPunct(tokens, "*") {
			node = newNode(ND_MUL)
		} else if isPunct(tokens, "/") {
			node = newNode(ND_DIV)
		} else if isPunct(tokens, "%") {
			node = newNode(ND_MOD)
		} else {
			break
		}
		tokens = tokens[1:]
		node.Lhs = lhs
		rhs, rest := unaryExpr(tokens)
		node.Rhs = rhs
		tokens = rest
		lhs = node
	}
	return lhs, tokens
}

func unaryExpr(tokens []*tokenizer.Token) (*Node, []*tokenizer.Token) {
	if isPunct(tokens, "-") {
		tokens = tokens[1:]
		node := newNode(ND_UNARY_MINUS)
		lhs, rest := primaryExpr(tokens)
		node.Lhs = lhs
		return node, rest
	}
	return primaryExpr(tokens)
}

func primaryExpr(tokens []*tokenizer.Token) (*Node, []*tokenizer.Token) {
	if isNum(tokens) {
		node := newNode(ND_NUM)
		node.Num = tokens[0].Num
		return node, tokens[1:]
	}
	if isPunct(tokens, "(") {
		tokens = tokens[1:]
		node, rest := expr(tokens)
		if !isPunct(rest, ")") {
			panic("unmatched open bracket")
		}
		return node, rest[1:]
	}
	panic("invalid expression")
}

func isPunct(tokens []*tokenizer.Token, punct string) bool {
	return len(tokens) > 0 && tokens[0].Kind == tokenizer.TK_PUNCT && bytes.Equal(tokens[0].Text, []byte(punct))
}

func isNum(tokens []*tokenizer.Token) bool {
	return len(tokens) > 0 && tokens[0].Kind == tokenizer.TK_NUM
}

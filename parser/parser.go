package parser

import (
	"bytes"

	"github.com/Div9851/mini-go/tokenizer"
)

type Parser struct {
	Tokens    []*tokenizer.Token
	Rest      []*tokenizer.Token
	Offset    map[string]int
	StackSize int
}

func New(tokens []*tokenizer.Token) *Parser {
	parser := new(Parser)
	parser.Tokens = tokens
	parser.Rest = parser.Tokens
	parser.Offset = make(map[string]int)
	parser.StackSize = 0
	return parser
}

func (p *Parser) Parse() []*Node {
	var nodes []*Node
	for p.Rest[0].Kind != tokenizer.TK_EOF {
		node := p.simpleStmt()
		nodes = append(nodes, node)
	}
	return nodes
}

func (p *Parser) simpleStmt() *Node {
	tok := p.Rest[0]
	if matchPunct(p.Rest[0], ";") {
		p.Rest = p.Rest[1:]
		return newNode(ND_EMPTY_STMT, tok)
	}
	lhs := p.exprList()
	if matchPunct(p.Rest[0], ";") {
		p.Rest = p.Rest[1:]
		if len(lhs.Exprs) > 1 {
			panic("invalid expression statement")
		}
		node := newNode(ND_EXPR_STMT, tok)
		node.Lhs = lhs.Exprs[0]
		return node
	}
	if matchPunct(p.Rest[0], "=") {
		p.Rest = p.Rest[1:]
		rhs := p.exprList()
		if !matchPunct(p.Rest[0], ";") {
			panic("missing semicolon")
		}
		p.Rest = p.Rest[1:]
		node := newNode(ND_ASSIGN_STMT, tok)
		node.Lhs = lhs
		node.Rhs = rhs
		return node
	}
	if matchPunct(p.Rest[0], ":=") {
		for _, expr := range lhs.Exprs {
			if expr.Kind != ND_VAR {
				panic("invalid statement(var)")
			}
			p.StackSize += 8
			p.Offset[expr.VarName] = p.StackSize
			expr.Offset = p.StackSize
		}
		p.Rest = p.Rest[1:]
		rhs := p.exprList()
		if !matchPunct(p.Rest[0], ";") {
			panic("missing semicolon")
		}
		p.Rest = p.Rest[1:]
		node := newNode(ND_ASSIGN_STMT, tok)
		node.Lhs = lhs
		node.Rhs = rhs
		return node
	}
	panic("invalid statement")
}

func (p *Parser) exprList() *Node {
	tok := p.Rest[0]
	node := newNode(ND_EXPR_LIST, tok)
	for {
		node.Exprs = append(node.Exprs, p.expr())
		if !matchPunct(p.Rest[0], ",") {
			break
		}
		p.Rest = p.Rest[1:]
	}
	return node
}

func (p *Parser) expr() *Node {
	return p.addExpr()
}

func (p *Parser) addExpr() *Node {
	tok := p.Rest[0]
	lhs := p.mulExpr()
	for {
		var node *Node
		if matchPunct(p.Rest[0], "+") {
			node = newNode(ND_ADD, tok)
		} else if matchPunct(p.Rest[0], "-") {
			node = newNode(ND_SUB, tok)
		} else {
			break
		}
		p.Rest = p.Rest[1:]
		node.Lhs = lhs
		rhs := p.mulExpr()
		node.Rhs = rhs
		lhs = node
	}
	return lhs
}

func (p *Parser) mulExpr() *Node {
	tok := p.Rest[0]
	lhs := p.unaryExpr()
	for {
		var node *Node
		if matchPunct(p.Rest[0], "*") {
			node = newNode(ND_MUL, tok)
		} else if matchPunct(p.Rest[0], "/") {
			node = newNode(ND_DIV, tok)
		} else if matchPunct(p.Rest[0], "%") {
			node = newNode(ND_MOD, tok)
		} else {
			break
		}
		p.Rest = p.Rest[1:]
		node.Lhs = lhs
		rhs := p.unaryExpr()
		node.Rhs = rhs
		lhs = node
	}
	return lhs
}

func (p *Parser) unaryExpr() *Node {
	tok := p.Rest[0]
	if matchPunct(p.Rest[0], "-") {
		p.Rest = p.Rest[1:]
		node := newNode(ND_UNARY_MINUS, tok)
		lhs := p.primaryExpr()
		node.Lhs = lhs
		return node
	}
	return p.primaryExpr()
}

func (p *Parser) primaryExpr() *Node {
	tok := p.Rest[0]
	if matchNum(p.Rest[0]) {
		node := newNode(ND_NUM, tok)
		node.Num = p.Rest[0].Num
		p.Rest = p.Rest[1:]
		return node
	}
	if matchIdent(p.Rest[0]) {
		node := newNode(ND_VAR, tok)
		node.VarName = string(p.Rest[0].Text)
		offset, ok := p.Offset[node.VarName]
		if ok {
			node.Offset = offset
		} else {
			node.Offset = -1
		}
		p.Rest = p.Rest[1:]
		return node
	}
	if matchPunct(p.Rest[0], "(") {
		p.Rest = p.Rest[1:]
		node := p.expr()
		if !matchPunct(p.Rest[0], ")") {
			panic("unmatched open bracket")
		}
		p.Rest = p.Rest[1:]
		return node
	}
	panic("invalid expression")
}

func matchPunct(token *tokenizer.Token, punct string) bool {
	return token.Kind == tokenizer.TK_PUNCT && bytes.Equal(token.Text, []byte(punct))
}

func matchNum(token *tokenizer.Token) bool {
	return token.Kind == tokenizer.TK_NUM
}

func matchIdent(token *tokenizer.Token) bool {
	return token.Kind == tokenizer.TK_IDENT
}

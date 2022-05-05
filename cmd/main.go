package main

import (
	"os"

	"github.com/Div9851/mini-go/codegen"
	"github.com/Div9851/mini-go/parser"
	"github.com/Div9851/mini-go/tokenizer"
)

func main() {
	input := []byte(os.Args[1])
	t := tokenizer.New(input)
	tokens := t.Tokenize()
	p := parser.New(tokens)
	nodes := p.Parse()
	codegen.Gen(nodes, p.StackSize)
}

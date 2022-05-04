package main

import (
	"os"

	"github.com/Div9851/mini-go/codegen"
	"github.com/Div9851/mini-go/parser"
	"github.com/Div9851/mini-go/tokenizer"
)

func main() {
	input := []byte(os.Args[1])
	tokens := tokenizer.Tokenize(input)
	node := parser.Parse(tokens)
	codegen.Generate(node)
}

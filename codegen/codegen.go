package codegen

import (
	"fmt"

	"github.com/Div9851/mini-go/parser"
)

func Generate(node *parser.Node) {
	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".globl main")
	fmt.Println("main:")
	generateExpr(node)
	fmt.Println("    pop rax")
	fmt.Println("    ret")
}

func generateExpr(node *parser.Node) {
	if node.Kind == parser.ND_NUM {
		fmt.Printf("    push %d\n", node.Num)
		return
	}
	generateExpr(node.Lhs)
	generateExpr(node.Rhs)
	fmt.Println("    pop rdi")
	fmt.Println("    pop rax")
	switch node.Kind {
	case parser.ND_ADD:
		fmt.Println("    add rax, rdi")
	case parser.ND_SUB:
		fmt.Println("    sub rax, rdi")
	case parser.ND_MUL:
		fmt.Println("    imul rax, rdi")
	case parser.ND_DIV:
		fmt.Println("    cqo")
		fmt.Println("    idiv rdi")
	case parser.ND_MOD:
		fmt.Println("    cqo")
		fmt.Println("    idiv rdi")
		fmt.Println("    mov rax, rdx")
	default:
		panic("invalid expression")
	}
	fmt.Println("    push rax")
}

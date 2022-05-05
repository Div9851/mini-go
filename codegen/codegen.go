package codegen

import (
	"fmt"

	"github.com/Div9851/mini-go/parser"
)

func Gen(nodes []*parser.Node, stackSize int) {
	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".globl main")
	fmt.Println("main:")
	fmt.Println("    push rbp")
	fmt.Println("    mov rbp, rsp")
	fmt.Printf("    sub rsp, %d\n", stackSize)
	for _, node := range nodes {
		genStmt(node)
	}
	fmt.Println("    mov rsp, rbp")
	fmt.Println("    pop rbp")
	fmt.Println("    ret")
}

func genLval(node *parser.Node) {
	if node.Kind != parser.ND_VAR || node.Offset == -1 {
		panic("invalid lval")
	}
	fmt.Println("    mov rax, rbp")
	fmt.Printf("    sub rax, %d\n", node.Offset)
}

func genStmt(node *parser.Node) {
	switch node.Kind {
	case parser.ND_EMPTY_STMT:
	case parser.ND_EXPR_STMT:
		genExpr(node.Lhs)
	case parser.ND_ASSIGN_STMT:
		rhsExprs := node.Rhs.Exprs
		for i := len(rhsExprs) - 1; i >= 0; i-- {
			expr := rhsExprs[i]
			genExpr(expr)
			fmt.Println("    push rax")
		}
		lhsExprs := node.Lhs.Exprs
		for _, expr := range lhsExprs {
			genLval(expr)
			fmt.Println("    pop rdi")
			fmt.Println("    mov [rax], rdi")
		}
	}
}

func genExpr(node *parser.Node) {
	if node.Kind == parser.ND_NUM {
		fmt.Printf("    mov rax, %d\n", node.Num)
		return
	}
	if node.Kind == parser.ND_VAR {
		genLval(node)
		fmt.Println("    mov rax, [rax]")
		return
	}
	genExpr(node.Rhs)
	fmt.Println("    push rax")
	genExpr(node.Lhs)
	fmt.Println("    pop rdi")
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
}

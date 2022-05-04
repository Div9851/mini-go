package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := strconv.Atoi(os.Args[1])
	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".globl main")
	fmt.Println("main:")
	fmt.Printf("	mov rax, %d\n", input)
	fmt.Println("	ret")
}

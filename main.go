package main

import (
	"fmt"
	"gulu-interpreter/lexer"
)

func main() {
	var lex *lexer.Lexer
	lex = lexer.New("12344")

	fmt.Println(lex.NextToken())
	// lexer.Lexer(c)
	fmt.Println("hello")
}

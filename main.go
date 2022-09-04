package main

import (
	"fmt"
	"gulu-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the gulu programming language! \n",
		user.Username)
	fmt.Printf("Feel free to type in commads\n")
	repl.Start(os.Stdin, os.Stdout)
}

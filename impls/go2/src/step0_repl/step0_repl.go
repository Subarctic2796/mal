package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func READ(src string) string { return src }

func EVAL(ast string, env string) string {
	_ = env
	return ast
}

func PRINT(exp string) string { return exp }

func rep(src string) string { return PRINT(EVAL(READ(src), "")) }

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("user> ")
		if !stdin.Scan() {
			break
		}

		src := strings.TrimSpace(stdin.Text())
		fmt.Println(rep(src))
	}
}

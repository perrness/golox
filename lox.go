package main

import (
	"bufio"
	"fmt"
	"loxgo/lox"
	"os"
)

func main() {
	args := os.Args

	if len(args) > 2 {
		println("Usage: jlox [script]")
		os.Exit(64)
	} else if len(args) == 2 {
		runFile(args[1])
	} else {
		runPrompt()
	}
}

func runFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Error reading file: %v\n", err))
	}

	hadError := run(string(bytes))

	if hadError {
		os.Exit(65)
	}
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	print("> ")
	for scanner.Scan() {
		line := scanner.Text()

		run(line)
		print("> ")
	}
}

func run(source string) bool {
	scanner := lox.NewScanner(source)

	tokens := scanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}

	return false
}

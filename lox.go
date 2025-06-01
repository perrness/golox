package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	run(string(bytes))
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

func run(source string) {
	scanner := bufio.NewScanner(strings.NewReader(source))
	scanner.Split(bufio.ScanWords)

	var tokens []string

	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}

	for _, token := range tokens {
		fmt.Println(token)
	}
}

package lox

import (
	"fmt"
	"os"
)

var hadError = false

func error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf("[line %d] Error%s: %s", line, where, message))
	hadError = true
}

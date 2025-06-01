package lox

import (
	"fmt"
	"os"
)

func error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) bool {
	fmt.Fprintln(os.Stderr, fmt.Sprintf("[line %d] Error%s: %s", line, where, message))
	return true
}

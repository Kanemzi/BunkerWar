package util

import (
	tm "github.com/buger/goterm"
)

// OutPlayerName puts the string in yellow bold between parenthesis
func OutPlayerName(name string) string {
	return "[" + tm.Color(tm.Bold(name), tm.YELLOW) + "]"
}

// OutBaseName puts the string in green bold between parenthesis
func OutBaseName(name string) string {
	return "[" + tm.Color(tm.Bold(name), tm.GREEN) + "]"
}

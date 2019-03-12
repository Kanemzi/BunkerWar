package main

import (
	tm "github.com/buger/goterm"
)

// out_PlayerName puts the string in yellow bold between parenthesis
func outPlayerName(name string) string {
	return "[" + tm.Color(tm.Bold(name), tm.YELLOW) + "]"
}

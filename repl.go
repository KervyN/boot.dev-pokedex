package main

import (
	"strings"
)

func cleanInput(input string) []string {
	input = strings.Trim(input, " ")
	input = strings.ToLower(input)
	return strings.Split(input, " ")
}

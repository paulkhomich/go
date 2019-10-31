package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "$t Hardware or software? So hard to choose. $c1 Software. $c2 Hardware."

	println(s)
	println(expand(s, f))
}

func expand(s string, f func(string) string) string {
	words := strings.Split(s, " ")

	for i, v := range words {
		if v[0] == '$' {
			words[i] = f(v[1:])
		}
	}

	return strings.Join(words, " ")
}

func f(s string) string {
	if s == "t" {
		return "[Title]"
	} else if s[0] == 'c' {
		return fmt.Sprintf("[Chapter %s]", s[1:])
	}

	return ""
}

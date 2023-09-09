package main

import "strings"

func parseInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	return strings.Split(text, " ")
}
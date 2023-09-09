package main

import (
	"bufio"
	"os"
)

func main() {

	reader := bufio.NewScanner(os.Stdin)
	runner := NewRunner()
	for reader.Scan() {
		input := parseInput(reader.Text())
		runner.exeCommand(input[0])
	}
}
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewScanner(os.Stdin)
	runner := NewRunner()
	fmt.Print("Pokedex > ")
	for reader.Scan() {
		input := parseInput(reader.Text())
		runner.exeCommand(input[0])
		fmt.Print("Pokedex > ")
	}
}
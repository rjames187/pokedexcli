package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		fmt.Print(reader.Text())
	}
}
package main

import (
	"fmt"
	"os"
	pokeapi "pokedexcli/internal"
)

type cliCommand struct {
	name string
	description string
	callback func(map[string]cliCommand) error
	config *pokeapi.PageConfig
}

type commandRunner struct {
	commands map[string]cliCommand
}

func (c *commandRunner) exeCommand(name string) {
	command, ok := c.commands[name]
	if !ok {
		fmt.Printf("Error! %s is not a command", name)
		os.Exit(1)
	}
	err := command.callback(c.commands)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func NewRunner() *commandRunner {
	runner := &commandRunner{commands: map[string]cliCommand{}}
	runner.commands["help"] = cliCommand{
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	}
	runner.commands["exit"] = cliCommand{
		name: "exit",
		description: "exits the Pokedex",
		callback: commandExit,
	}
	pageConfig := &pokeapi.PageConfig{}
	runner.commands["map"] = cliCommand{
		name: "map",
		description: "explores the map going forwards",
		callback: commandMap,
		config: pageConfig,
	}
	runner.commands["mapb"] = cliCommand{
		name: "mapb",
		description: "explores the map going backwards",
		callback: commandMapb,
		config: pageConfig,
	}
	return runner
}

func commandHelp(commands map[string]cliCommand) error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}

func commandExit(commands map[string]cliCommand) error {
	os.Exit(0)
	return nil
}

func commandMap(commands map[string]cliCommand) error {
	mapCommand := commands["map"]
	locations, config, err := pokeapi.GetLocations(mapCommand.config, true)
	if err != nil {
		return err
	}
	mapCommand.config = config
	commands["map"] = mapCommand
	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}

func commandMapb(commands map[string]cliCommand) error {
	mapCommand := commands["mapb"]
	locations, config, err := pokeapi.GetLocations(mapCommand.config, false)
	if err != nil {
		return err
	}
	mapCommand.config = config
	commands["mapb"] = mapCommand
	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}
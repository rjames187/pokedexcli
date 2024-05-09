package main

import (
	"fmt"
	"os"
	pokeapi "pokedexcli/internal"
)

type cliCommand struct {
	name string
	description string
	callback func(*commandConfig) error
}

type commandRunner struct {
	config *commandConfig
}

type commandConfig struct {
	commands map[string]cliCommand
	pageConfig *pokeapi.PageConfig
}

func (c *commandRunner) exeCommand(name string) {
	command, ok := c.config.commands[name]
	if !ok {
		fmt.Printf("Error! %s is not a command", name)
		os.Exit(1)
	}
	err := command.callback(c.config)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func NewRunner() *commandRunner {
	config := &commandConfig{map[string]cliCommand{}, nil}
	config.commands["help"] = cliCommand{
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	}
	config.commands["exit"] = cliCommand{
		name: "exit",
		description: "exits the Pokedex",
		callback: commandExit,
	}
	config.commands["map"] = cliCommand{
		name: "map",
		description: "explores the map going forwards",
		callback: commandMap,
	}
	config.commands["mapb"] = cliCommand{
		name: "mapb",
		description: "explores the map going backwards",
		callback: commandMapb,
	}
	config.pageConfig = &pokeapi.PageConfig{}
	return &commandRunner{config}
}

func commandHelp(config *commandConfig) error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range config.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}

func commandExit(config *commandConfig) error {
	os.Exit(0)
	return nil
}

func commandMap(config *commandConfig) error {
	locations, newPageConfig, err := pokeapi.GetLocations(config.pageConfig, true)
	if err != nil {
		return err
	}
	config.pageConfig = newPageConfig
	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}

func commandMapb(config *commandConfig) error {
	locations, newPageConfig, err := pokeapi.GetLocations(config.pageConfig, false)
	if err != nil {
		return err
	}
	config.pageConfig = newPageConfig
	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}
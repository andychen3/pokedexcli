package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Commands struct {
	name        string
	description string
	callback    func()
}

func getCommands() map[string]Commands {
	return map[string]Commands{
		"help": {
			name:        "help",
			description: "Displays the help menu.",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex.",
			callback:    commandExit,
		},
	}
}

func commandHelp() {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
}

func commandExit() {
	os.Exit(0)

}

func cleanInput(s string) []string {
	output := strings.Split(strings.ToLower(s), " ")
	return output
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command. Try again.")
		} else {
			command.callback()
		}

	}
}

package engine

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartGame() error {
	fmt.Println("Welcome to Cairn Console RPG!")
	fmt.Println("Type 'help' for available commands.")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "help":
			fmt.Println("Available commands: help, look, quit")
		case "look":
			fmt.Println("You stand in a dark wood. All is quiet.")
		case "quit":
			fmt.Println("Farewell, traveler.")
			return nil
		default:
			fmt.Println("Unrecognized command. Try 'help'.")
		}
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/NightMarethyu/GoCairnCLI/engine"
)

func main() {
	err := engine.StartGame()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error starting game: ", err)
		os.Exit(1)
	}
}

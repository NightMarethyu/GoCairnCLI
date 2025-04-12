package engine

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/NightMarethyu/go_console_rpg/game"
	"github.com/NightMarethyu/go_console_rpg/player"
	"github.com/NightMarethyu/go_console_rpg/world"
)

func convertToPointerSlice(items []game.Item) []*game.Item {
	pointerSlice := make([]*game.Item, len(items))
	for i := range items {
		pointerSlice[i] = &items[i]
	}
	return pointerSlice
}

func StartGame() error {
	backgrounds, err := game.LoadBackgrounds()
	if err != nil {
		return err
	}

	var bg = backgrounds[0]

	itemMap, err := game.LoadItems()
	if err != nil {
		return err
	}

	startingItems, err := game.ResolveStartingItems(bg, itemMap)
	if err != nil {
		return err
	}

	enemies, err := game.LoadEnemies()
	if err != nil {
		return err
	}

	var enemy = enemies[0]

	ctx := &GameContext{
		inCombat:   false,
		Background: &bg,
		Inventory:  convertToPointerSlice(startingItems),
		Player:     &player.Player{Name: "Hero", HP: 10},
		Enemy:      &enemy,
		Location: &world.Location{
			Name:        "Starting Area",
			Description: "You are in a dark forest. The trees are tall and the path is overgrown.",
		},
	}

	ctx.inCombat = true

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nWhat would you like to do?")
		actions := GetAvailableActions(ctx)

		for i, act := range actions {
			fmt.Printf(" [%d] %s - %s\n", i+1, act.Label, act.Description)
		}

		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		matched := false
		for i, act := range actions {
			if input == act.Label || input == fmt.Sprint(i+1) {
				act.Execute(ctx)
				matched = true
				break
			}
		}

		if !matched {
			fmt.Println("Invalid action. Please try again.")

		}
	}
}

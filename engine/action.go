package engine

import (
	"fmt"
	"os"

	"github.com/NightMarethyu/go_console_rpg/game"
)

type Action struct {
	Label       string
	Description string
	Execute     func(*GameContext)
	Condition   func() bool
}

func GetAvailableActions(ctx *GameContext) []Action {
	actions := []Action{}

	if ctx.inCombat {
		actions = append(actions, Action{
			Label:       "Attack",
			Description: "Attack with your equipped weapon.",
			Execute: func(c *GameContext) {
				fmt.Printf("You attack the %s!\n", c.Enemy.Name)
				c.Enemy.HP -= 3 // Replace with the weapon damage dice value
				if c.Enemy.HP <= 0 {
					fmt.Printf("You defeated %s!", c.Enemy.Name)
					c.inCombat = false
					c.Enemy = nil
				}
			},
		})
	} else {
		actions = append(actions, Action{
			Label:       "look",
			Description: "Survey your surroundings.",
			Execute: func(c *GameContext) {
				fmt.Printf("%s", c.Location.Description)
			},
		})
	}

	if ctx.PlayerHasItem("healing_potion") {
		actions = append(actions, Action{
			Label:       "drink potion",
			Description: "Use a healing potion to restore HP",
			Execute: func(c *GameContext) {
				fmt.Println("You drink a healing potion")
				c.Player.HP += 5 // Replace with the potion healing value
				c.Inventory = removeItem(c.Inventory, "healing_potion")
			},
		})
	}

	actions = append(actions, Action{
		Label:       "inventory",
		Description: "Check your inventory",
		Execute: func(c *GameContext) {
			fmt.Println("Your inventory contains:")
			for _, item := range c.Inventory {
				fmt.Printf("- %s", item.Name)
			}
		},
	})

	actions = append(actions, Action{
		Label:       "quit",
		Description: "Exit the game",
		Execute: func(c *GameContext) {
			fmt.Println("Thanks for playing!")
			os.Exit(0)
		},
	})

	return actions
}

func removeItem(items []*game.Item, id string) []*game.Item {
	result := make([]*game.Item, 0, len(items))
	removed := false
	for _, item := range items {
		if item.ID == id && !removed {
			removed = true
			continue
		}
		result = append(result, item)
	}
	return result
}

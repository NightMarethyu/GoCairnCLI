package engine

import (
	"github.com/NightMarethyu/go_console_rpg/game"
	"github.com/NightMarethyu/go_console_rpg/player"
	"github.com/NightMarethyu/go_console_rpg/world"
)

type GameContext struct {
	inCombat   bool
	Inventory  []*game.Item
	Player     *player.Player
	Enemy      *game.Enemy
	Background *game.Background
	Location   *world.Location
}

func (ctx *GameContext) PlayerHasItem(id string) bool {
	for _, item := range ctx.Inventory {
		if item.ID == id {
			return true
		}
	}
	return false
}

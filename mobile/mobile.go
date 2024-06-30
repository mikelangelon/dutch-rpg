package mobile

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"
	"github.com/mikelangelon/dutchrpg/game"
)

func init() {
	mobile.SetGame(game.NewGame())
}

func Dummy() {}

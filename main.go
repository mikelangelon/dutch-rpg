package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/dutchrpg/assets"
	"github.com/mikelangelon/dutchrpg/game"
	"github.com/mikelangelon/dutchrpg/graphics"
	"log/slog"
)

func main() {
	initialMap, err := graphics.NewMapScene(assets.MapPackPNG, assets.InitialMap, assets.MapPackTSX, 1000, 600, 3)
	if err != nil {
		slog.Error("crash parseTileSet", "error", err)
		return
	}
	factory, _ := graphics.NewCharFactory(assets.MapPackPNG, assets.MapPackTSX, 3)
	playerImage := factory.CharImage(361)
	player := &graphics.Char{
		ID:     "player",
		Image:  playerImage,
		X:      0 * 16,
		Y:      7 * 16,
		ScaleX: 3,
		ScaleY: 3,
	}
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game.NewGame(initialMap, player)); err != nil {
		slog.With("error", err).Error("unexpected error running game")
	}
}

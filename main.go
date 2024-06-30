package main

import (
	"github.com/mikelangelon/dutchrpg/game"
	"log/slog"
)

func main() {
	if err := ebiten.RunGame(&game.Game{}); err != nil {
		slog.With("error", err).Error("unexpected error running game")
	}
}

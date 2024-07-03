package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/dutchrpg/assets"
	"github.com/mikelangelon/dutchrpg/core"
	"github.com/mikelangelon/dutchrpg/game"
	"gopkg.in/yaml.v3"
	"log/slog"
)

func main() {
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game.NewGame(parseWords())); err != nil {
		slog.With("error", err).Error("unexpected error running game")
	}
}

func parseWords() []*core.Word {
	var words []*core.Word
	err := yaml.Unmarshal(assets.Nouns, &words)
	if err != nil {
		slog.Error("error unmarshalling words", "error", err)
	}
	return words
}

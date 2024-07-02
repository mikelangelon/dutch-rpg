package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/dutchrpg/ui"
)

const (
	ScreenWidth  = 720
	ScreenHeight = 1280
)

type Game struct {
	UI displayer
}

func NewGame() *Game {
	return &Game{UI: ui.NewQuestionsUI()}
}

func (g *Game) Update() error {
	return g.UI.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.UI.Draw(screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return ScreenWidth, ScreenHeight
}

type displayer interface {
	Update() error
	Draw(screen *ebiten.Image)
}

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
	Menu *ui.MenuUI
}

func NewGame() *Game {
	return &Game{Menu: ui.NewMenuUI()}
}

func (g *Game) Update() error {
	g.Menu.UI.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Menu.UI.Draw(screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return ScreenWidth, ScreenHeight
}

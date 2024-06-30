package game

import "github.com/hajimehoshi/ebiten/v2"

const (
	ScreenWidth  = 1280
	ScreenHeight = 720
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {}

func (g *Game) Layout(_, _ int) (int, int) {
	return ScreenWidth, ScreenHeight
}

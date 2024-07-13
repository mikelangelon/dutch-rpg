package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/town-sweet-town/common"
)

const frameSpeed = 20

type Char struct {
	ID            string
	Image         *ebiten.Image
	IdleAnimation []*ebiten.Image
	X, Y          int64

	ScaleX, ScaleY float64
	frameCount     int
}

func (c *Char) Position() common.Position {
	return common.Position{X: c.X, Y: c.Y}
}

func (c *Char) Update() {
	c.frameCount++
}
func (c *Char) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.X), float64(c.Y))
	op.GeoM.Scale(c.ScaleX, c.ScaleY)

	image := c.IdleAnimation[(c.frameCount/frameSpeed)%len(c.IdleAnimation)]
	screen.DrawImage(image, op)
}

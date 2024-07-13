package scene

import (
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
)

type HUI struct {
	ui ebitenui.UI

	Hearth  *ebiten.Image
	hearths int
}

func NewHUI(Hearth *ebiten.Image, hearths int) *HUI {
	return &HUI{Hearth: Hearth, hearths: hearths}
}

func (h *HUI) Update() error {
	return nil
}

func (h *HUI) Hearths(hearths int) {
	h.hearths = hearths
}

func (h *HUI) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(340), float64(75))
	op.GeoM.Scale(2, 2)
	for range h.hearths {
		screen.DrawImage(h.Hearth, op)
		op.GeoM.Translate(-40, 0)
	}
}

package ui

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"image/color"
)

type MenuUI struct {
	UI *ebitenui.UI
}

var menuOptions = []string{"Study"}

func NewMenuUI() *MenuUI {
	container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{true}),
			widget.GridLayoutOpts.Spacing(20, 5),
			widget.GridLayoutOpts.Padding(widget.Insets{Left: 200, Right: 200, Top: 200, Bottom: 200}),
		),
		),
	)
	buttonsContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{true, true, true, true}),
			widget.GridLayoutOpts.Spacing(20, 5),
		)),
	)
	container.AddChild(buttonsContainer)
	for _, v := range menuOptions {
		b := createButton(v)
		buttonsContainer.AddChild(b)
	}
	return &MenuUI{UI: &ebitenui.UI{
		Container: container,
	}}
}

func createButton(optionText string) *widget.Button {
	buttonImage, _ := loadButtonImage()

	face, _ := loadFont(40)
	return widget.NewButton(
		widget.ButtonOpts.Image(buttonImage),
		widget.ButtonOpts.Text(optionText, face, &widget.ButtonTextColor{
			Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
		}),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   200,
			Right:  200,
			Top:    10,
			Bottom: 10,
		}),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
			})),
	)
}

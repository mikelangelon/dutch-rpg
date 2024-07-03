package ui

import (
	"fmt"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"image/color"
)

func loadButtonImage() (*widget.ButtonImage, error) {
	idle := image.NewNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255})
	hover := image.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})
	pressed := image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}

func loadFont(size float64) (font.Face, error) {
	ttfFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, fmt.Errorf("problem parsing font: %w", err)
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}

type buttonFunc func(args *widget.ButtonClickedEventArgs)

func createButton(optionText string, f buttonFunc) *widget.Button {
	buttonImage, _ := loadButtonImage()

	face, _ := loadFont(40)
	return widget.NewButton(
		widget.ButtonOpts.Image(buttonImage),
		widget.ButtonOpts.Text(optionText, face, &widget.ButtonTextColor{
			Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
		}),
		widget.ButtonOpts.ClickedHandler(widget.ButtonClickedHandlerFunc(f)),
		//widget.ButtonOpts.TextPadding(widget.Insets{
		//	Left:   100,
		//	Right:  100,
		//	Top:    80,
		//	Bottom: 80,
		//}),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(300, 200),
		),
	)
}

package ui

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type QuestionsUI struct {
	UI *ebitenui.UI
}

func (q QuestionsUI) Update() error {
	q.UI.Update()
	return nil
}

func (q QuestionsUI) Draw(screen *ebiten.Image) {
	q.UI.Draw(screen)
}

func NewQuestionsUI() *QuestionsUI {
	face, _ := loadFont(20)
	container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0, 0, 0, 150})),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{true}),
			widget.GridLayoutOpts.Spacing(0, 0),
		),
		),
	)
	firstContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{100, 200, 250, 150})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(1, 200),
		),
	)
	firstContainer.AddChild(widget.NewText(
		widget.TextOpts.Text("Example", face, color.White),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	))

	secondContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{200, 0, 100, 150})),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(50),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top: 50,
			}),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
			widget.WidgetOpts.MinSize(100, 1280/2),
		),
	)

	//buttonsContainer := widget.NewContainer(
	//	widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0, 0, 0, 150})),
	//	widget.ContainerOpts.Layout(widget.NewGridLayout(
	//		widget.GridLayoutOpts.Columns(2),
	//		widget.GridLayoutOpts.Stretch([]bool{true, true}, []bool{true, true}),
	//		widget.GridLayoutOpts.Spacing(40, 40),
	//	),
	//	),
	//)

	//buttonsContainer := widget.NewContainer(
	//	widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{200, 0, 100, 150})),
	//	widget.ContainerOpts.Layout(widget.NewGridLayout()),
	//	widget.ContainerOpts.WidgetOpts(
	//		widget.WidgetOpts.MinSize(100, 1280/2),
	//		widget.WidgetOpts.LayoutData(widget.RowLayoutData{
	//			Position: widget.RowLayoutPositionCenter,
	//		}),
	//	),
	//)

	buttonsContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0, 0, 0, 150})),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),
			widget.GridLayoutOpts.Stretch([]bool{true, true}, []bool{true, true}),
			widget.GridLayoutOpts.Spacing(40, 40),
			widget.GridLayoutOpts.Padding(widget.Insets{
				Top:    50,
				Bottom: 20,
				Left:   50,
				Right:  50,
			}),
		),
		),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: widget.RowLayoutPositionCenter,
		})),
	)
	secondContainer.AddChild(widget.NewText(
		widget.TextOpts.Text("Some word in Dutch", face, color.White),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
		),
	))
	for _, v := range []string{"Option A", "Option B", "Option C", "Option D"} {
		b := createButton(v, func(args *widget.ButtonClickedEventArgs) {
			println(args.Button.Text().Label)
		})
		buttonsContainer.AddChild(b)
	}
	container.AddChild(firstContainer)
	container.AddChild(secondContainer)
	secondContainer.AddChild(buttonsContainer)
	return &QuestionsUI{
		UI: &ebitenui.UI{Container: container},
	}
}

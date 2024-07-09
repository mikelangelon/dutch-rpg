package ui

import (
	"fmt"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/dutchrpg/core"
	"image/color"
)

type QuestionsUI struct {
	UI *ebitenui.UI
	// HUI
	buttons     []*widget.Button
	wordLabel   *widget.Text
	pointsLabel *widget.Text
	// To return back
	answer *string
}

func (q *QuestionsUI) Update() error {
	q.UI.Update()
	return nil
}

func (q *QuestionsUI) Draw(screen *ebiten.Image) {
	q.UI.Draw(screen)
}

func (q *QuestionsUI) SetQuestion(question core.Question, points int) {
	q.answer = nil
	for i, v := range q.buttons {
		v.Text().Label = question.Options[i]
	}
	q.wordLabel.Label = question.Word
	q.pointsLabel.Label = fmt.Sprintf("%d", points)
}

func (q *QuestionsUI) GetAnswer() *string {
	return q.answer
}

func NewQuestionsUI() *QuestionsUI {
	face, _ := loadFont(40)
	pointsFont, _ := loadFont(100)
	container := widget.NewContainer(

		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{true}),
			widget.GridLayoutOpts.Spacing(0, 0),
		),
		),
	)
	ui := &QuestionsUI{
		UI: &ebitenui.UI{Container: container},
	}
	firstContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(50),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top: 300,
			}),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
		),
	)
	pointsLabel := widget.NewText(
		widget.TextOpts.Text("Example", pointsFont, color.White),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),
	)
	ui.pointsLabel = pointsLabel
	firstContainer.AddChild(pointsLabel)

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
	wordLabel := widget.NewText(
		widget.TextOpts.Text("Some word in Dutch", face, color.White),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
		),
	)
	ui.wordLabel = wordLabel
	secondContainer.AddChild(wordLabel)

	for _, v := range []string{"", "", "", ""} {
		b := createButton(v, func(args *widget.ButtonClickedEventArgs) {
			ui.answer = &args.Button.Text().Label
		})
		ui.buttons = append(ui.buttons, b)
		buttonsContainer.AddChild(b)
	}
	container.AddChild(firstContainer)
	container.AddChild(secondContainer)
	secondContainer.AddChild(buttonsContainer)
	return ui
}

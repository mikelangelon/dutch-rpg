package ui

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/dutchrpg/core"
	"image/color"
)

type BasisUI struct {
	UI *ebitenui.UI
	// HUI
	mainLabel        *widget.Text
	challenger       challenger
	buttonsContainer *widget.Container
}

type challenger interface {
	SetQuestion(question core.Question, container *widget.Container)
	GetAnswer() *string
}

func (q *BasisUI) Update() error {
	q.UI.Update()
	return nil
}

func (q *BasisUI) Draw(screen *ebiten.Image) {
	q.UI.Draw(screen)
}

func (q *BasisUI) GetAnswer() *string {
	return q.challenger.GetAnswer()
}
func (q *BasisUI) SetQuestion(question core.Question) {
	q.buttonsContainer.RemoveChildren()
	q.challenger.SetQuestion(question, q.buttonsContainer)
	q.mainLabel.Label = question.Word
}
func NewBasisUI() *BasisUI {
	face, _ := loadFont(40)
	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{true}),
			widget.GridLayoutOpts.Spacing(0, 0),
		),
		),
	)
	ui := &BasisUI{
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
	ui.mainLabel = wordLabel
	secondContainer.AddChild(wordLabel)
	container.AddChild(firstContainer)
	container.AddChild(secondContainer)
	secondContainer.AddChild(buttonsContainer)
	ui.buttonsContainer = buttonsContainer
	ui.challenger = &questionsChallenge{}
	return ui
}

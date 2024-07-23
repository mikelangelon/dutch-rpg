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
	secondaryLabel   *widget.Text
	challenger       challenger
	buttonsContainer *widget.Container
	labelsContainer  *widget.Container
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
	val, ok := q.challenger.(*spellingChallenge)
	if ok {
		q.secondaryLabel.Label = val.builtAnswer
	}
	switch q.challenger.(type) {
	case *spellingChallenge:
		val, _ := q.challenger.(*spellingChallenge)
		q.secondaryLabel.Label = val.builtAnswer
	case *orderChallenge:
		val, _ := q.challenger.(*orderChallenge)
		q.secondaryLabel.Label = val.builtAnswer
	}
	return q.challenger.GetAnswer()
}

func (q *BasisUI) SetQuestion(question core.Question) {
	q.UI.Container.RemoveChild(q.labelsContainer)
	switch question.Type {
	case "questions":
		q.challenger = &questionsChallenge{}
		q.createQuestionsUI(question, 50, 50, 2)
	case "yes-no":
		q.challenger = &yesnoChallenge{}
		q.createQuestionsUI(question, 10, 10, 2)
	case "spelling":
		q.challenger = &spellingChallenge{}
		q.createQuestionsUI(question, 10, 10, 3)
	case "order":
		q.challenger = &orderChallenge{}
		q.createQuestionsUI(question, 10, 10, 3)
	}
	q.challenger.SetQuestion(question, q.buttonsContainer)
}

func (q *BasisUI) createQuestionsUI(question core.Question, spacing, top, columns int) {
	face, _ := loadFont(40)
	q.labelsContainer = widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{200, 0, 100, 150})),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(spacing),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top: top,
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
			widget.GridLayoutOpts.Columns(columns),
			widget.GridLayoutOpts.Stretch([]bool{true, true, true}, []bool{true, true, true}),
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
		widget.TextOpts.Text(question.Word, face, color.White),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
		),
	)
	q.mainLabel = wordLabel
	q.labelsContainer.AddChild(wordLabel)
	if question.SecondaryWord != nil {
		secondaryLabel := widget.NewText(
			widget.TextOpts.Text(*question.SecondaryWord, face, color.White),
			widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
			widget.TextOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.RowLayoutData{
					Position: widget.RowLayoutPositionCenter,
					Stretch:  true,
				}),
			),
		)
		q.secondaryLabel = secondaryLabel
		q.labelsContainer.AddChild(secondaryLabel)
	}
	q.labelsContainer.AddChild(buttonsContainer)
	q.buttonsContainer = buttonsContainer
	q.UI.Container.AddChild(q.labelsContainer)
}

func NewBasisUI() *BasisUI {
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
	container.AddChild(firstContainer)
	ui.createQuestionsUI(core.Question{}, 0, 0, 2)
	return ui
}

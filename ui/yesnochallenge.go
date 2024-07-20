package ui

import (
	"github.com/ebitenui/ebitenui/widget"
	"github.com/mikelangelon/dutchrpg/core"
)

type yesnoChallenge struct {
	buttons  []*widget.Button
	question string
	// To return back
	answer *string
}

func (q *yesnoChallenge) SetQuestion(question core.Question, container *widget.Container) {
	q.answer = nil
	for _, v := range question.Options {
		b := createButtonOpts(v, func(args *widget.ButtonClickedEventArgs) {
			q.answer = &args.Button.Text().Label
		}, widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(300, 400),
		))
		if len(v) < 15 {
			b.Text().Face, _ = loadFont(40)
		} else if len(v) < 23 {
			b.Text().Face, _ = loadFont(30)
		} else {
			b.Text().Face, _ = loadFont(20)
		}
		q.buttons = append(q.buttons, b)
		container.AddChild(b)
	}
	q.question = question.Word
}

func (q *yesnoChallenge) GetAnswer() *string {
	return q.answer
}

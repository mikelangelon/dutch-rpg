package ui

import (
	"fmt"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/mikelangelon/dutchrpg/core"
	"strings"
)

type orderChallenge struct {
	buttons  []*widget.Button
	question string
	// To return back
	answer      *string
	builtAnswer string
}

func (q *orderChallenge) SetQuestion(question core.Question, container *widget.Container) {
	q.answer = nil
	q.builtAnswer = ""
	for _, v := range question.Options {
		b := createButtonOpts(v, func(args *widget.ButtonClickedEventArgs) {
			if q.builtAnswer == "" {
				q.builtAnswer = args.Button.Text().Label
			} else {
				q.builtAnswer = fmt.Sprintf("%s %s", q.builtAnswer, args.Button.Text().Label)
			}
			if !strings.Contains(question.Answer, q.builtAnswer) {
				q.answer = &q.builtAnswer
			} else if q.builtAnswer == question.Answer {
				q.answer = &q.builtAnswer
			}
		}, widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(150, 70),
		))
		q.buttons = append(q.buttons, b)
		b.Text().Face, _ = loadFont(20)
		container.AddChild(b)
	}
	q.question = question.Word
}

func (q *orderChallenge) GetAnswer() *string {
	return q.answer
}

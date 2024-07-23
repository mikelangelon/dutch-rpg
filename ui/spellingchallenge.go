package ui

import (
	"fmt"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/mikelangelon/dutchrpg/core"
	"math/rand"
	"strings"
)

type spellingChallenge struct {
	buttons  []*widget.Button
	question string
	// To return back
	answer      *string
	builtAnswer string
}

func (q *spellingChallenge) SetQuestion(question core.Question, container *widget.Container) {
	q.answer = nil
	q.builtAnswer = ""

	for _, v := range question.Options {
		b := createButtonOpts(v, func(args *widget.ButtonClickedEventArgs) {
			q.builtAnswer = fmt.Sprintf("%s%s", q.builtAnswer, args.Button.Text().Label)
			if !strings.Contains(question.Answer, q.builtAnswer) {
				q.answer = &q.builtAnswer
			} else if q.builtAnswer == question.Answer {
				q.answer = &q.builtAnswer
			}
		}, widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(100, 100),
		))
		q.buttons = append(q.buttons, b)
		container.AddChild(b)
	}
	q.question = question.Word
}

func (q *spellingChallenge) GetAnswer() *string {
	return q.answer
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

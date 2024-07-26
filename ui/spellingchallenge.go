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
	q.builtAnswer = " "
	options := randomizeOptions(string(question.Answer[0]), 8)
	for _, v := range options {
		b := createButtonOpts(v, func(args *widget.ButtonClickedEventArgs) {
			if len(strings.Trim(q.builtAnswer, " ")) == 0 {
				q.builtAnswer = ""
			}
			q.builtAnswer = fmt.Sprintf("%s%s", q.builtAnswer, args.Button.Text().Label)
			if !strings.Contains(question.Answer, q.builtAnswer) {
				q.answer = &q.builtAnswer
			} else if q.builtAnswer == question.Answer {
				q.answer = &q.builtAnswer
			} else {
				nextStep := strings.Replace(question.Answer, q.builtAnswer, "", 1)
				options = randomizeOptions(string(nextStep[0]), 8)
				q.SetOptions(options)
			}
		}, widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(100, 100),
		))
		q.buttons = append(q.buttons, b)
		container.AddChild(b)
	}
	q.question = question.Word
}

func (q *spellingChallenge) SetOptions(options []string) {
	for i, v := range q.buttons {
		v.Text().Label = options[i]
	}
}
func (q *spellingChallenge) GetAnswer() *string {
	return q.answer
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ ")

func randLetter(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func randomizeOptions(correct string, letters int) []string {
	result := []string{correct}
	for i := 0; i < letters; i++ {
		result = append(result, string(letterRunes[rand.Intn(len(letterRunes))]))
	}
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}

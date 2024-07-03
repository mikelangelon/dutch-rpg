package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/dutchrpg/core"
	"github.com/mikelangelon/dutchrpg/ui"
	"math/rand"
)

const (
	ScreenWidth  = 720
	ScreenHeight = 1280
)

const (
	statusNextWord = "NEXT_QUESTION"
	statusWaiting  = "WAITING"
)

type Game struct {
	UI              displayer
	Words           []*core.Word
	Status          string
	CounterCorrect  int
	currentQuestion core.Question
}

func NewGame(words []*core.Word) *Game {
	return &Game{
		UI:     ui.NewQuestionsUI(),
		Words:  words,
		Status: statusNextWord,
	}
}

func (g *Game) randomWord() *core.Word {
	return g.Words[rand.Intn(len(g.Words))]
}
func (g *Game) prepareQuestion() core.Question {
	w := g.randomWord()
	option1 := g.randomWord().English
	option2 := g.randomWord().English
	option3 := g.randomWord().English
	options := []string{w.English, option1, option2, option3}
	rand.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })
	return core.Question{
		Word:    w.Dutch,
		Answer:  w.English,
		Options: options,
	}
}
func (g *Game) Update() error {
	switch g.Status {
	case statusNextWord:
		g.currentQuestion = g.prepareQuestion()
		g.UI.SetQuestion(g.currentQuestion, g.CounterCorrect)
		g.Status = statusWaiting
	case statusWaiting:
		if answer := g.UI.GetAnswer(); answer != nil {
			g.Status = statusNextWord
			if g.currentQuestion.Answer == *answer {
				g.CounterCorrect++
			} else {
				g.CounterCorrect = 0
			}

		}
	}
	return g.UI.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.UI.Draw(screen)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return ScreenWidth, ScreenHeight
}

type displayer interface {
	Update() error
	Draw(screen *ebiten.Image)
	SetQuestion(question core.Question, points int)
	GetAnswer() *string
}

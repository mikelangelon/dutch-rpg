package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/dutchrpg/assets"
	"github.com/mikelangelon/dutchrpg/core"
	"github.com/mikelangelon/dutchrpg/graphics"
	"github.com/mikelangelon/dutchrpg/ui"
	"gopkg.in/yaml.v3"
	"log/slog"
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

	Scene  *graphics.MapScene
	Player *graphics.Char
}

func NewGame(scene *graphics.MapScene, player *graphics.Char) *Game {
	return &Game{
		UI:     ui.NewQuestionsUI(),
		Words:  parseWords(),
		Status: statusNextWord,
		Scene:  scene,
		Player: player,
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
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.Scene.Camera.Position[0] += 1
	}
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
				g.Player.X += 16
				if g.Player.X > 14*16 {
					g.Scene.Camera.Position[0] += 14 * 3 * 16
					g.Player.X = 1 * 16
				}
			} else {
				g.CounterCorrect = 0
			}

		}
	}
	return g.UI.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.UI.Draw(screen)
	g.Scene.Draw(screen)
	g.Player.Draw(screen)
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

func parseWords() []*core.Word {
	var words []*core.Word
	err := yaml.Unmarshal(assets.Nouns, &words)
	if err != nil {
		slog.Error("error unmarshalling words", "error", err)
	}
	return words
}

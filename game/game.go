package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/dutchrpg/assets"
	"github.com/mikelangelon/dutchrpg/core"
	"github.com/mikelangelon/dutchrpg/graphics"
	"github.com/mikelangelon/dutchrpg/graphics/scene"
	"github.com/mikelangelon/dutchrpg/persistence"
	"github.com/mikelangelon/dutchrpg/ui"
	"log/slog"
	"math/rand"
)

const (
	ScreenWidth  = 720
	ScreenHeight = 1280
	initialLives = 3
)

const (
	statusNextWord = "NEXT_QUESTION"
	statusWaiting  = "WAITING"
)

type Game struct {
	UI              displayer
	WordsStore      *persistence.WordStore
	Status          string
	CounterCorrect  int
	lives           int
	currentQuestion core.Question

	HUI    *scene.HUI
	Scene  *graphics.MapScene
	Player *graphics.Char
}

func NewGame() *Game {
	initialMap, err := graphics.NewMapScene(assets.MapPackPNG, assets.InitialMap, assets.MapPackTSX, 50*16*3, 600, 3)
	if err != nil {
		slog.Error("crash parseTileSet", "error", err)
		return nil
	}
	factory, _ := graphics.NewCharFactory(assets.MapPackPNG, assets.MapPackTSX, 3)
	playerImage := factory.CharImage(361)
	hearth := factory.CharImage(532)
	player := &graphics.Char{
		ID:            "player",
		Image:         playerImage,
		IdleAnimation: []*ebiten.Image{factory.CharImage(361), factory.CharImage(363), factory.CharImage(365)},
		X:             1 * 16,
		Y:             7 * 16,
		ScaleX:        3,
		ScaleY:        3,
	}
	return &Game{
		UI:         ui.NewQuestionsUI(),
		WordsStore: persistence.New(),
		Status:     statusNextWord,
		Scene:      initialMap,
		Player:     player,
		lives:      initialLives,
		HUI:        scene.NewHUI(hearth, initialLives),
	}
}

func (g *Game) prepareQuestion() core.Question {
	difficulty := g.CounterCorrect / 2
	if difficulty > 8 {
		difficulty = 8
	}
	w := g.WordsStore.WordDifficulty(difficulty)
	option1 := g.WordsStore.RandomWord().English
	option2 := g.WordsStore.RandomWord().English
	option3 := g.WordsStore.RandomWord().English
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
				g.lives -= 1
				if g.lives <= 0 {
					// gameover. For now, starting again
					g.lives = initialLives
					g.Scene.Camera.Position[0] = 0
					g.Player.X = 1 * 16
				}
				g.HUI.Hearths(g.lives)
			}

		}
	}
	g.Player.Update()
	return g.UI.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.UI.Draw(screen)
	g.Scene.Draw(screen)
	g.Player.Draw(screen)
	g.HUI.Draw(screen)
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

package persistence

import (
	"github.com/mikelangelon/dutchrpg/assets"
	"github.com/mikelangelon/dutchrpg/core"
	"gopkg.in/yaml.v3"
	"log/slog"
	"math/rand"
)

type WordStore struct {
	Words []*core.Word

	Flexibility int
}

func New() *WordStore {
	ws := &WordStore{Words: parseWords()}
	ws.Shuffle()
	ws.Flexibility = 1
	return ws
}

func (ws *WordStore) RandomWord() *core.Word {
	return ws.Words[rand.Intn(len(ws.Words))]
}

func (ws *WordStore) Shuffle() {
	rand.Shuffle(len(ws.Words), func(i, j int) {
		ws.Words[i], ws.Words[j] = ws.Words[j], ws.Words[i]
	})
}
func (ws *WordStore) WordDifficulty(difficulty int) *core.Word {
	for i, v := range ws.Words {
		if v.Difficulty >= difficulty-ws.Flexibility && v.Difficulty <= difficulty+ws.Flexibility {
			ws.Words = append(ws.Words[0:i], append(ws.Words[i+1:], ws.Words[i])...)
			return v
		}
	}
	return nil
}

func parseWords() []*core.Word {
	var words []*core.Word
	err := yaml.Unmarshal(assets.Nouns, &words)
	if err != nil {
		slog.Error("error unmarshalling words", "error", err)
	}
	return words
}

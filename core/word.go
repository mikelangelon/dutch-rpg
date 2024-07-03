package core

type Word struct {
	Dutch     string   `yaml:"dutch"`
	English   string   `yaml:"english"`
	Type      string   `yaml:"type"`
	Sentences []string `yaml:"sentences"`
}

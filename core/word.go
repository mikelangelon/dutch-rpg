package core

type Word struct {
	Dutch      string   `yaml:"dutch"`
	English    string   `yaml:"english"`
	Type       string   `yaml:"type"`
	DeHet      *string  `yaml:"dehet,omitempty"`
	Difficulty int      `yaml:"difficulty"`
	Sentences  []string `yaml:"sentences"`
	Labels     []string `yaml:"labels"`
}

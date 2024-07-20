package core

type Question struct {
	Word          string
	SecondaryWord *string
	Answer        string
	Options       []string
	Type          string
}

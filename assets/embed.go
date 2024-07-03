package assets

import (
	_ "embed"
)

var (
	//go:embed 1000nouns.yaml
	Nouns []byte
)

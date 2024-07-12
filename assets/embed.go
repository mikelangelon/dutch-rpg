package assets

import (
	_ "embed"
)

var (
	//go:embed 1000nouns.yaml
	Nouns []byte

	//go:embed colored_packed.tsx
	MapPackTSX []byte

	//go:embed colored_packed.png
	MapPackPNG []byte

	//go:embed initialMap.tmx
	InitialMap []byte
)

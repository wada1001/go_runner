package cmp

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/yohamta/donburi"
)

type (
	TitleData struct {
		Title  string
		Cursor int
	}

	FontData struct {
		Face *text.GoTextFaceSource
	}
)

var (
	Title = donburi.NewComponentType[TitleData]()
	Font  = donburi.NewComponentType[FontData]()
)

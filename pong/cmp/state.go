package cmp

import "github.com/yohamta/donburi"

type (
	StateData struct {
		Current, Next string
	}
)

var (
	State = donburi.NewComponentType[StateData]()
)

const (
	TitleState = "title"
	GameState  = "game"
)

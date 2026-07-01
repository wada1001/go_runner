package cmp

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type (
	PongSettingsData struct {
		BallSpeed                    float64
		BallX, BallY                 float64
		BallRadius                   float64
		PaddleWidth                  float64
		PaddleX, PaddleY             float64
		PaddleSpeed                  float64
		PaddleHeight                 float64
		WallLeft, WallRight, WallTop float64
	}

	ScoreData struct {
		Score int
	}

	InputData struct {
		PressedKeys      []ebiten.Key
		JustPressedKeys  []ebiten.Key
		JustReleasedKeys []ebiten.Key
	}

	WallsData struct {
		Left, Right, Top float64
	}

	PositionData struct {
		X, Y float64
	}

	VelocityData struct {
		X, Y float64
	}

	CircleColliderData struct {
		Radius float64
	}

	RectColliderData struct {
		Width, Height float64
	}
)

var (
	BallTag   = donburi.NewTag("ball")
	PaddleTag = donburi.NewTag("paddle")
	WallsTag  = donburi.NewTag("walls")

	Input = donburi.NewComponentType[InputData]()

	Score          = donburi.NewComponentType[ScoreData]()
	PongSettings   = donburi.NewComponentType[PongSettingsData]()
	Walls          = donburi.NewComponentType[WallsData]()
	Position       = donburi.NewComponentType[PositionData]()
	Velocity       = donburi.NewComponentType[VelocityData]()
	CircleCollider = donburi.NewComponentType[CircleColliderData]()
	RectCollider   = donburi.NewComponentType[RectColliderData]()
)

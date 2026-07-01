package cmp

import (
	"github.com/yohamta/donburi"
)

type (
	CommandsData struct {
		Commands []Command
	}

	Command interface {
		Exec(w donburi.World) error
	}

	CreatePongCmd struct{}

	CreateTitleCmd struct{}
)

var (
	Commands = donburi.NewComponentType[CommandsData]()
)

func (c *CreateTitleCmd) Exec(w donburi.World) error {
	entry := w.Entry(w.Create(Title))
	Title.SetValue(entry, TitleData{
		Title: "Pong",
	})

	return nil
}

func (c *CreatePongCmd) Exec(w donburi.World) error {
	settings := PongSettings.Get(PongSettings.MustFirst(w))

	w.Create(Score)

	ballEntry := w.Entry(w.Create(BallTag, Position, Velocity, CircleCollider))
	Position.Set(ballEntry, &PositionData{X: settings.BallX, Y: settings.BallY})
	Velocity.Set(ballEntry, &VelocityData{X: settings.BallSpeed, Y: settings.BallSpeed})
	CircleCollider.Set(ballEntry, &CircleColliderData{Radius: settings.BallRadius})

	wallsEntry := w.Entry(w.Create(WallsTag, Walls))
	Walls.SetValue(wallsEntry, WallsData{
		Left:  settings.WallLeft,
		Right: settings.WallRight,
		Top:   settings.WallTop,
	})

	paddleEntry := w.Entry(w.Create(PaddleTag, Position, Velocity, RectCollider))
	Position.Set(paddleEntry, &PositionData{X: settings.PaddleX, Y: settings.PaddleY})
	RectCollider.Set(paddleEntry, &RectColliderData{Width: settings.PaddleWidth, Height: settings.PaddleHeight})

	return nil
}

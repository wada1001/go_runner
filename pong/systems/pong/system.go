package pong

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
	"github.com/wada1001/go_runner/pong/cmp"
	"github.com/yohamta/donburi/ecs"
)

func InputGame(e *ecs.ECS) {
	paddleEntry := cmp.PaddleTag.MustFirst(e.World)
	vel := cmp.Velocity.Get(paddleEntry)
	vel.X = 0

	inp := cmp.Input.Get(cmp.Input.MustFirst(e.World))
	if len(inp.PressedKeys) == 0 {
		return
	}

	settings := cmp.PongSettings.Get(cmp.PongSettings.MustFirst(e.World))
	if lo.Contains(inp.PressedKeys, ebiten.KeyArrowLeft) {
		vel.X = -settings.PaddleSpeed
		return
	}

	if lo.Contains(inp.PressedKeys, ebiten.KeyArrowRight) {
		vel.X = settings.PaddleSpeed
		return
	}

	return
}

func UpdateBall(e *ecs.ECS) {
	ballEntry := cmp.BallTag.MustFirst(e.World)

	pos := cmp.Position.Get(ballEntry)
	vel := cmp.Velocity.Get(ballEntry)

	pos.X += vel.X
	pos.Y += vel.Y
	return
}

func UpdatePaddle(e *ecs.ECS) {
	paddleEntry := cmp.PaddleTag.MustFirst(e.World)

	pos := cmp.Position.Get(paddleEntry)
	vel := cmp.Velocity.Get(paddleEntry)

	pos.X += vel.X
	return
}

func CollidePaddle(e *ecs.ECS) {
	walls := cmp.Walls.Get(cmp.Walls.MustFirst(e.World))

	paddleEntry := cmp.PaddleTag.MustFirst(e.World)
	pos := cmp.Position.Get(paddleEntry)
	if pos.X < walls.Left {
		pos.X = walls.Left
		return
	}

	rect := cmp.RectCollider.Get(paddleEntry)
	if pos.X+rect.Width > walls.Right {
		pos.X = walls.Right - rect.Width
	}

	return
}

func BounceWalls(e *ecs.ECS) {
	walls := cmp.Walls.Get(cmp.Walls.MustFirst(e.World))

	ballEntry := cmp.BallTag.MustFirst(e.World)
	pos := cmp.Position.Get(ballEntry)
	vel := cmp.Velocity.Get(ballEntry)
	collider := cmp.CircleCollider.Get(ballEntry)

	if pos.X-collider.Radius < walls.Left {
		vel.X = -vel.X
		pos.X = walls.Left + collider.Radius
	}

	if pos.X+collider.Radius > walls.Right {
		vel.X = -vel.X
		pos.X = walls.Right - collider.Radius
	}

	if pos.Y-collider.Radius < walls.Top {
		vel.Y = -vel.Y
	}

	return
}

func BouncePaddle(e *ecs.ECS) {
	paddleEntry := cmp.PaddleTag.MustFirst(e.World)
	paddlePos := cmp.Position.Get(paddleEntry)
	paddleRect := cmp.RectCollider.Get(paddleEntry)

	ballEntry := cmp.BallTag.MustFirst(e.World)
	ballPos := cmp.Position.Get(ballEntry)
	if ballPos.Y < paddlePos.Y || ballPos.Y > paddlePos.Y+paddleRect.Height {
		return
	}

	ballVel := cmp.Velocity.Get(ballEntry)
	ballCol := cmp.CircleCollider.Get(ballEntry)
	if ballPos.X+ballCol.Radius < paddlePos.X || ballPos.X-ballCol.Radius > paddlePos.X+paddleRect.Width {
		return
	}

	score := cmp.Score.Get(cmp.Score.MustFirst(e.World))
	score.Score += 1

	paddleCenter := paddlePos.X + paddleRect.Width/2

	offset := (ballPos.X - paddleCenter) / (paddleRect.Width / 2)

	if offset < -1 {
		offset = -1
	} else if offset > 1 {
		offset = 1
	}

	maxAngle := math.Pi / 3
	angle := offset * maxAngle
	speed := math.Hypot(ballVel.X, ballVel.Y)
	ballVel.X = speed * math.Sin(angle)
	ballVel.Y = -speed * math.Cos(angle)

	minY := 2.0
	if math.Abs(ballVel.Y) < minY {
		if ballVel.Y < 0 {
			ballVel.Y = -minY
		} else {
			ballVel.Y = minY
		}
	}

	ballPos.Y = paddlePos.Y - ballCol.Radius
	return
}

func ControlGameDifficulty(e *ecs.ECS) {
	score := cmp.Score.Get(cmp.Score.MustFirst(e.World))
	settings := cmp.PongSettings.Get(cmp.PongSettings.MustFirst(e.World))

	ballEntry := cmp.BallTag.MustFirst(e.World)
	ballVel := cmp.Velocity.Get(ballEntry)

	speed := settings.BallSpeed + float64(score.Score)*0.2

	ballVel.X = lo.Ternary(ballVel.X > 0, speed, -speed)
	ballVel.Y = lo.Ternary(ballVel.Y > 0, speed, -speed)
}

func ObserveGameEnd(e *ecs.ECS) {
	ballEntry := cmp.BallTag.MustFirst(e.World)

	pos := cmp.Position.Get(ballEntry)
	if pos.Y < 600 {
		return
	}

	state := cmp.State.Get(cmp.State.MustFirst(e.World))
	state.Next = cmp.TitleState
}

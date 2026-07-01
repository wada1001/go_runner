package renderer

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/wada1001/go_runner/pong/cmp"
	"github.com/yohamta/donburi/ecs"
)

func RenderWalls(ecs *ecs.ECS, screen *ebiten.Image) {
	entry, ok := cmp.Walls.First(ecs.World)
	if !ok {
		return
	}

	walls := cmp.Walls.GetValue(entry)
	vector.StrokeLine(screen, float32(walls.Left), float32(walls.Top), float32(walls.Right), float32(walls.Top), 2, color.White, true)
	vector.StrokeLine(screen, float32(walls.Left), float32(walls.Top), float32(walls.Left), 600, 2, color.White, true)
	vector.StrokeLine(screen, float32(walls.Right), float32(walls.Top), float32(walls.Right), 600, 2, color.White, true)
}

func RenderBall(ecs *ecs.ECS, screen *ebiten.Image) {
	entry, ok := cmp.BallTag.First(ecs.World)
	if !ok {
		return
	}

	pos := cmp.Position.GetValue(entry)
	col := cmp.CircleCollider.GetValue(entry)
	vector.StrokeCircle(screen, float32(pos.X), float32(pos.Y), float32(col.Radius), 2, color.White, true)
}

func RenderPaddle(ecs *ecs.ECS, screen *ebiten.Image) {
	entry, ok := cmp.PaddleTag.First(ecs.World)
	if !ok {
		return
	}

	pos := cmp.Position.GetValue(entry)
	rect := cmp.RectCollider.GetValue(entry)
	vector.StrokeLine(screen, float32(pos.X), float32(pos.Y), float32(pos.X+rect.Width), float32(pos.Y), 2, color.White, true)
}

func RenderScore(ecs *ecs.ECS, screen *ebiten.Image) {
	entry, ok := cmp.Score.First(ecs.World)
	if !ok {
		return
	}

	score := cmp.Score.GetValue(entry)

	font := cmp.Font.Get(cmp.Font.MustFirst(ecs.World))
	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(float64(30), float64(10))
	textOp.ColorScale.ScaleWithColor(color.White)
	textOp.SecondaryAlign = text.AlignCenter

	text.Draw(screen, fmt.Sprintf("Score: %d", score.Score), &text.GoTextFace{
		Source: font.Face,
		Size:   15,
	}, textOp)
}

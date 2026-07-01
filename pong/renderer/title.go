package renderer

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/wada1001/go_runner/pong/cmp"
	"github.com/yohamta/donburi/ecs"
)

func RenderTitle(ecs *ecs.ECS, screen *ebiten.Image) {
	titleEntry, ok := cmp.Title.First(ecs.World)
	if !ok {
		return
	}

	font := cmp.Font.Get(cmp.Font.MustFirst(ecs.World))

	// Draw title
	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(float64(200), float64(200))
	textOp.ColorScale.ScaleWithColor(color.White)
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter
	title := cmp.Title.GetValue(titleEntry)
	text.Draw(screen, title.Title, &text.GoTextFace{
		Source: font.Face,
		Size:   50,
	}, textOp)

	textOp.GeoM.Reset()
	textOp.GeoM.Translate(float64(200), float64(300))
	textOp.ColorScale.ScaleWithColor(color.White)
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter
	text.Draw(screen, "Press Space to Start", &text.GoTextFace{
		Source: font.Face,
		Size:   20,
	}, textOp)
}

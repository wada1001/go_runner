package input

import (
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/wada1001/go_runner/pong/cmp"
	"github.com/yohamta/donburi/ecs"
)

func Update(e *ecs.ECS) {
	inp := cmp.Input.Get(cmp.Input.MustFirst(e.World))
	inp.PressedKeys = inpututil.AppendPressedKeys(inp.PressedKeys[:0])
	inp.JustPressedKeys = inpututil.AppendJustPressedKeys(inp.JustPressedKeys[:0])
	inp.JustReleasedKeys = inpututil.AppendJustReleasedKeys(inp.JustReleasedKeys[:0])
}

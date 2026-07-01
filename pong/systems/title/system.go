package title

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
	"github.com/wada1001/go_runner/pong/cmp"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

func InputTitle(e *ecs.ECS) {
	inp := cmp.Input.Get(cmp.Input.MustFirst(e.World))

	if len(inp.JustReleasedKeys) == 0 {
		return
	}

	title := cmp.Title.Get(cmp.Title.MustFirst(e.World))
	if lo.Contains(inp.JustReleasedKeys, ebiten.KeyArrowUp) {
		title.Cursor = lo.Ternary(title.Cursor > 0, title.Cursor-1, 0)
		return
	}

	if lo.Contains(inp.JustReleasedKeys, ebiten.KeyArrowDown) {
		title.Cursor = lo.Ternary(title.Cursor >= 3, 3, title.Cursor+1)
		return
	}

	if !lo.Contains(inp.JustReleasedKeys, ebiten.KeySpace) {
		return
	}
	state := cmp.State.Get(cmp.State.MustFirst(e.World))
	state.Next = cmp.GameState
}

func ClearEntities(e *ecs.ECS) {
	query := donburi.NewQuery(filter.Not(filter.Contains(cmp.GlobalTag)))
	for entry := range query.Iter(e.World) {
		e.World.Remove(entry.Entity())
	}
}

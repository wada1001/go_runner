package observestate

import (
	"github.com/wada1001/go_runner/pong/cmp"
	"github.com/yohamta/donburi/ecs"
)

func Update(e *ecs.ECS) {
	state := cmp.State.Get(cmp.State.MustFirst(e.World))
	if state.Current == state.Next {
		return
	}

	state.Current = state.Next
	return
}

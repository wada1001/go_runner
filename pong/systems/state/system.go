package state

import (
	"github.com/wada1001/go_runner/pong/cmp"
	"github.com/yohamta/donburi/ecs"
)

func RunIf(target string, system ecs.System) ecs.System {
	return func(e *ecs.ECS) {
		state := cmp.State.Get(cmp.State.MustFirst(e.World))
		if state.Current != target {
			return
		}

		system(e)
	}
}

func Enter(target string, system ecs.System) ecs.System {
	return func(e *ecs.ECS) {
		state := cmp.State.Get(cmp.State.MustFirst(e.World))
		if state.Current == state.Next {
			return
		}

		if state.Next != target {
			return
		}

		system(e)
	}
}

func Exit(target string, system ecs.System) ecs.System {
	return func(e *ecs.ECS) {
		state := cmp.State.Get(cmp.State.MustFirst(e.World))
		if state.Current == state.Next {
			return
		}

		if state.Current != target {
			return
		}

		system(e)
	}
}

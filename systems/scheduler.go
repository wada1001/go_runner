package systems

import (
	"sort"

	"github.com/yohamta/donburi"
)

type (
	System    func(w donburi.World) error
	Condition func(w donburi.World) bool
	Systems   []System

	SystemSet struct {
		System System
		Order  int
	}

	Scheduler []SystemSet

	RunIf struct {
		system System
		cond   Condition
	}
)

func CreateScheduler(systems ...SystemSet) Scheduler {
	sort.SliceStable(systems, func(i, j int) bool {
		return systems[i].Order < systems[j].Order
	})

	return Scheduler(systems)
}

func (s Scheduler) Update(w donburi.World) {
	for _, system := range s {
		system.System(w)
	}
}

func CreateRunIf(cond Condition, system System) *RunIf {
	return &RunIf{
		cond:   cond,
		system: system,
	}
}

func (r *RunIf) Update(w donburi.World) error {
	if !r.cond(w) {
		return nil
	}

	return r.system(w)
}

func (s Systems) Update(w donburi.World) error {
	for _, system := range s {
		if err := system(w); err != nil {
			return err
		}
	}

	return nil
}

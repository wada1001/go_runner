package command

import (
	"github.com/wada1001/go_runner/pong/cmp"
	"github.com/yohamta/donburi/ecs"
)

func Update(e *ecs.ECS) {
	commands := cmp.Commands.Get(cmp.Commands.MustFirst(e.World))
	for _, cmd := range commands.Commands {
		if err := cmd.Exec(e.World); err != nil {
			panic(err)
		}
	}

	commands.Commands = commands.Commands[:0]
}

func AddCommands(commands []cmp.Command) ecs.System {
	c := commands

	return func(e *ecs.ECS) {
		commands := cmp.Commands.Get(cmp.Commands.MustFirst(e.World))

		dst := make([]cmp.Command, len(c))
		copy(dst, c)

		commands.Commands = append(commands.Commands, dst...)
	}
}

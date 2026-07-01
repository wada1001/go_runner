package game

import (
	"bytes"
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/wada1001/go_runner/pong/cmp"
	"github.com/wada1001/go_runner/pong/renderer"
	"github.com/wada1001/go_runner/pong/systems/command"
	"github.com/wada1001/go_runner/pong/systems/input"
	"github.com/wada1001/go_runner/pong/systems/observestate"
	"github.com/wada1001/go_runner/pong/systems/pong"
	"github.com/wada1001/go_runner/pong/systems/state"
	"github.com/wada1001/go_runner/pong/systems/title"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

var (
	//go:embed noto.ttf
	font []byte
)

type (
	Game struct {
		ECS ecs.ECS
	}

	System func(w donburi.World) error
)

func Create() *Game {
	f, err := text.NewGoTextFaceSource(bytes.NewReader(font))
	if err != nil {
		log.Fatal(err)
	}

	world := donburi.NewWorld()
	entry := world.Entry(world.Create(
		cmp.GlobalTag,
		cmp.Commands, cmp.State,
		cmp.PongSettings, cmp.Input,
		cmp.Font,
	))

	cmp.Font.Set(entry, &cmp.FontData{
		Face: f,
	})
	cmp.State.Set(entry, &cmp.StateData{
		Next: cmp.TitleState,
	})
	cmp.PongSettings.Set(entry, &cmp.PongSettingsData{
		BallSpeed:    2,
		BallRadius:   4,
		BallX:        200,
		BallY:        300,
		PaddleSpeed:  2,
		PaddleWidth:  50,
		PaddleHeight: 2,
		PaddleX:      175,
		PaddleY:      550,
		WallLeft:     50,
		WallRight:    350,
		WallTop:      50,
	})

	return &Game{
		ECS: *ecs.NewECS(world).
			AddSystem(input.Update).
			AddSystem(state.RunIf(cmp.TitleState, title.InputTitle)).
			AddSystem(state.RunIf(cmp.GameState, pong.InputGame)).
			AddSystem(state.RunIf(cmp.GameState, pong.ControlGameDifficulty)).
			AddSystem(state.RunIf(cmp.GameState, pong.UpdateBall)).
			AddSystem(state.RunIf(cmp.GameState, pong.UpdatePaddle)).
			AddSystem(state.RunIf(cmp.GameState, pong.CollidePaddle)).
			AddSystem(state.RunIf(cmp.GameState, pong.BounceWalls)).
			AddSystem(state.RunIf(cmp.GameState, pong.BouncePaddle)).
			AddSystem(state.RunIf(cmp.GameState, pong.ObserveGameEnd)).
			// enter
			AddSystem(state.Enter(cmp.TitleState, command.AddCommands([]cmp.Command{
				&cmp.CreateTitleCmd{},
			}))).
			AddSystem(state.Enter(cmp.GameState, command.AddCommands([]cmp.Command{
				&cmp.CreatePongCmd{},
			}))).
			// exit
			AddSystem(state.Exit(cmp.TitleState, title.ClearEntities)).
			AddSystem(state.Exit(cmp.GameState, title.ClearEntities)).
			AddSystem(observestate.Update).
			AddSystem(command.Update).
			AddRenderer(ecs.LayerID(100), renderer.RenderTitle).
			AddRenderer(ecs.LayerID(100), renderer.RenderScore).
			AddRenderer(ecs.LayerID(5), renderer.RenderBall).
			AddRenderer(ecs.LayerID(5), renderer.RenderWalls).
			AddRenderer(ecs.LayerID(10), renderer.RenderPaddle),
	}
}

func (g *Game) Update() error {
	g.ECS.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ECS.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 400, 600
}

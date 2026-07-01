package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/wada1001/go_runner/pong/game"
)

func main() {
	ebiten.SetWindowSize(400, 600)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(game.Create()); err != nil {
		log.Fatal(err)
	}
}

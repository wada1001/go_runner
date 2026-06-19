package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/wada1001/go_runner/runner"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(runner.Create()); err != nil {
		log.Fatal(err)
	}
}

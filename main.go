// laberintogenerativo/main.go
package main

import (
	"laberintogenerativo/scenes"

	"github.com/hajimehoshi/ebiten"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

func main() {

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Laberinto Concurrente (Ebitengine Demo)")
	game, gameErr := scenes.NewGame()
	if gameErr != nil {
		panic(gameErr)
	}

	if runErr := game.Run(); runErr != nil {
		panic(runErr)
	}
}

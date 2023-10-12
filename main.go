// laberintogenerativo/main.go
package main

import (
	"laberintogenerativo/scenes"
)

func main() {
	game, gameErr := scenes.NewGame()
	if gameErr != nil {
		panic(gameErr)
	}

	if runErr := game.Run(); runErr != nil {
		panic(runErr)
	}
}

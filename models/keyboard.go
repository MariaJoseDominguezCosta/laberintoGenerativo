// laberintogenerativo/models/keyboard.go
package models

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func spacePressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySpace)
}

func SpaceReleased() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeySpace)
}

func upKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyUp)
}

func UpKeyReleased() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyUp)
}

func downKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyDown)
}

func DownKeyReleased() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyDown)
}

func leftKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyLeft)
}

func LeftKeyReleased() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyLeft)
}

func rightKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyRight)
}

func RightKeyReleased() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyRight)
}

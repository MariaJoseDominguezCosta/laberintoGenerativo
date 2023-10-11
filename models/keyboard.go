// laberintogenerativo/models/keyboard.go
package models

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func SpacePressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySpace)
}

func SpaceReleased() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeySpace)
}

func UpKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyUp)
}

func UpKeyReleased() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyUp)
}

func DownKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyDown)
}

func DownKeyReleased() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyDown)
}

func LeftKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyLeft)
}

func LeftKeyReleased() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyLeft)
}

func RightKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyRight)
}

func RightKeyReleased() bool {
	return inpututil.IsKeyJustReleased(ebiten.KeyRight)
}
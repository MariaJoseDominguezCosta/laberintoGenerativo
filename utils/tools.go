// laberintogenerativo/utils/tools.go
package utils

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

// Tile represents a space with an x,y coordinate within a Level. Any number of
// sprites may be added to a Tile.
type Tile struct {
	sprites []*ebiten.Image
}

// AddSprite adds a sprite to the Tile.
func (t *Tile) AddSprite(s *ebiten.Image) {
	t.sprites = append(t.sprites, s)
}

// Draw draws the Tile on the screen using the provided options.
func (t *Tile) Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions) {
	for _, s := range t.sprites {
		screen.DrawImage(s, options)
	}
}

// GetSprite returns a new image from source, of size width x height. Starting point of image is specified by the xoffset & yoffset.
func GetSprite(
	width, height int,
	xoffset, yoffset int,
	src *ebiten.Image,
) (*ebiten.Image, error) {
	sprite, spriteErr := ebiten.NewImage(width, height, ebiten.FilterDefault)
	if spriteErr != nil {
		return nil, spriteErr
	}

	rect := image.Rect(xoffset, yoffset, xoffset+width, yoffset+height)

	ops := &ebiten.DrawImageOptions{}
	ops.SourceRect = &rect
	if drawErr := sprite.DrawImage(src, ops); drawErr != nil {
		return nil, drawErr
	}

	return sprite, nil
}

// ScaleSprite returns a new image from source,
func ScaleSprite(src *ebiten.Image, x, y float64) (*ebiten.Image, error) {
	spriteW, spriteH := src.Size()
	sSprite, sSpriteErr := ebiten.NewImage(
		int(float64(spriteW)*x),
		int(float64(spriteH)*y),
		ebiten.FilterDefault)
	if sSpriteErr != nil {
		return nil, sSpriteErr
	}

	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Scale(x, y)
	if drawErr := sSprite.DrawImage(src, ops); drawErr != nil {
		return nil, drawErr
	}

	return sSprite, nil
}

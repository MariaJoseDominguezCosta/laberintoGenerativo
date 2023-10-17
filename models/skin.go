//laberintogenerativo/models/skin.go
package models
import (
	"image/color"
	"laberintogenerativo/resources"
	"laberintogenerativo/utils"
	"strconv"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
)

const MaxScoreView = 999999999
const MaxLifes = 7

func SkinView(
	skin *ebiten.Image,
	powers *resources.Powers,
	arcadeFont *truetype.Font,
) (func(state Mode, data *Data) (*ebiten.Image, error), error) {
	fontface := truetype.NewFace(arcadeFont, &truetype.Options{
		Size:    30,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	width, height := skin.Size()
	view, viewErr := ebiten.NewImage(width, height, ebiten.FilterDefault)
	if viewErr != nil {return nil, viewErr}

	life, lifeErr := utils.ScaleSprite(powers.Life, 0.5, 0.5)
	if lifeErr != nil {return nil, lifeErr}

	return func(state Mode, data *Data) (*ebiten.Image, error) {
		if clearErr := view.Clear(); clearErr != nil {
			return nil, clearErr
		}
		if drawErr := view.DrawImage(skin, &ebiten.DrawImageOptions{}); drawErr != nil {
			return nil, drawErr
		}
		switch state {
		case GameStart:
			fallthrough
		case GamePause:
			fallthrough
		case GameOver:
			if data != nil {
				score := data.Score
				lifes := data.Lifes

				if score > MaxScoreView {
					score = MaxScoreView
				}
				numstr := strconv.Itoa(score)
				text.Draw(view, numstr, fontface, 682-(len(numstr)*27), 64, color.RGBA{200, 150, 240,100})

				if lifes > MaxLifes {
					lifes = MaxLifes
				}

				ops := &ebiten.DrawImageOptions{}
				width, _ := life.Size()
				for i := 0; i < lifes; i++ {
					ops.GeoM.Reset()
					ops.GeoM.Translate(float64(650-(width*i)), 80)
					if drawErr := view.DrawImage(life, ops); drawErr != nil {
						return nil, drawErr
					}
				}
			}
		}

		return view, nil
	}, nil
}
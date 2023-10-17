// laberintogenerativo/models/grid.go
package models

import (
	"image/color"
	"laberintogenerativo/resources"
	"laberintogenerativo/utils"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

const GridViewSize = 1015

var GrayColor = color.RGBA{236, 240, 241, 255.0}

func GridView(
	characters *resources.Characters,
	powers *resources.Powers,
	arcadeFont *truetype.Font,
	mazeView func(state Mode, data *Data) (*ebiten.Image, error),
) (func(state Mode, data *Data) (*ebiten.Image, error), error) {
	fontface := truetype.NewFace(arcadeFont, &truetype.Options{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	limeAlpha := color.RGBA{250, 233, 8, 200}
	dot, dotErr := ebiten.NewImage(8, 8, ebiten.FilterDefault)
	if dotErr != nil {
		return nil, dotErr
	}
	if fillErr := dot.Fill(limeAlpha); fillErr != nil {
		return nil, fillErr
	}
	player, playerErr := utils.ScaleSprite(characters.Player, 0.5, 0.5)
	if playerErr != nil {
		return nil, playerErr
	}
	ghost1, ghost1Err := utils.ScaleSprite(characters.Ghost1, 0.5, 0.5)
	if ghost1Err != nil {
		return nil, ghost1Err
	}
	ghost2, ghost2Err := utils.ScaleSprite(characters.Ghost2, 0.5, 0.5)
	if ghost2Err != nil {
		return nil, ghost2Err
	}
	ghost3, ghost3Err := utils.ScaleSprite(characters.Ghost3, 0.5, 0.5)
	if ghost3Err != nil {
		return nil, ghost3Err
	}
	ghost4, ghost4Err := utils.ScaleSprite(characters.Ghost4, 0.5, 0.5)
	if ghost4Err != nil {
		return nil, ghost4Err
	}
	life, lifeErr := utils.ScaleSprite(powers.Life, 0.5, 0.5)
	if lifeErr != nil {
		return nil, lifeErr
	}
	invinci, invinciErr := utils.ScaleSprite(powers.Invincibility, 0.5, 0.5)
	if invinciErr != nil {
		return nil, invinciErr
	}
	view, viewErr := ebiten.NewImage(64*Columns, GridViewSize, ebiten.FilterDefault)
	if viewErr != nil {
		return nil, viewErr
	}
	return func(state Mode, data *Data) (*ebiten.Image, error) {
		if clearErr := view.Clear(); clearErr != nil {
			return nil, clearErr
		}
		if fillErr := view.Fill(color.RGBA{144, 144, 144, 255}); fillErr != nil {
			return nil, fillErr
		}
		ops := &ebiten.DrawImageOptions{}
		switch state {
		case GameLoading:
			text.Draw(view, "PRESS SPACE", fontface, 370-176, 512-(10+32), color.Black)
			text.Draw(view, "TO START", fontface, 370-128, 512+(10), color.Black)
		case GameStart, GamePause, GameOver:
			mazeView, mazeViewErr := mazeView(state, data)
			if mazeViewErr != nil {
				return nil, mazeViewErr
			}
			ops.GeoM.Reset()
			ops.GeoM.Translate(0, -(float64(len(data.Grid)*CellSize) - (GridViewSize + data.GridOffsetY)))
			if drawErr := view.DrawImage(mazeView, ops); drawErr != nil {
				return nil, drawErr
			}
			for i := 0; i < len(data.Active); i++ {
				for j := 0; j < Columns; j++ {
					if !data.Active[i][j] {
						ops.GeoM.Reset()
						ops.GeoM.Translate(float64((j*CellSize)+30),-(float64(((i*CellSize)+(CellSize/2))+2) - (GridViewSize + data.GridOffsetY)))
						if drawErr := view.DrawImage(dot, ops); drawErr != nil {
							return nil, drawErr
						}
					}
				}
			}
			for i := 0; i < len(data.Powers); i++ {
				power := data.Powers[i]
				powerImg := life
				if power.Kind == Invincibility {
					powerImg = invinci
				}
				pwidth, pheight := powerImg.Size()
				ops.GeoM.Reset()
				ops.GeoM.Translate(float64((data.Powers[i].Position.CellX*CellSize)+pwidth/2), -(float64(((data.Powers[i].Position.CellY*CellSize)+(CellSize/2))+pheight/2) - (GridViewSize + data.GridOffsetY)))
				if drawErr := view.DrawImage(powerImg, ops); drawErr != nil {
					return nil, drawErr
				}
			}
			ops.GeoM.Reset()
			pwidth, pheight := player.Size()
			switch data.Player.Position.Direction {
			case North:
				ops.GeoM.Rotate(-1.5708)
				ops.GeoM.Translate( data.Player.Position.PosX-float64(pwidth/2), GridViewSize-(data.Player.Position.PosY-float64(pheight-(pheight/2))))
			case East:
				ops.GeoM.Translate( data.Player.Position.PosX-float64(pwidth/2), GridViewSize-(data.Player.Position.PosY+float64(pheight/2)))
			case South:
				ops.GeoM.Rotate(1.5708)
				ops.GeoM.Translate( data.Player.Position.PosX+float64(pwidth/2), GridViewSize-(data.Player.Position.PosY+float64(pheight/2)))
			case West:
				ops.GeoM.Rotate(3.14159)
				ops.GeoM.Translate( data.Player.Position.PosX+float64(pwidth/2), GridViewSize-(data.Player.Position.PosY-float64(pheight-(pheight/2))))
			}
			if drawErr := view.DrawImage(player, ops); drawErr != nil {
				return nil, drawErr
			}
			for i := 0; i < len(data.Ghosts); i++ {
				ghost := data.Ghosts[i]
				ghostImg := ghost1
				switch ghost.Kind {
				case Ghost2:
					ghostImg = ghost2
				case Ghost3:
					ghostImg = ghost3
				case Ghost4:
					ghostImg = ghost4
				}
				gwidth, gheight := ghostImg.Size()
				ops.GeoM.Reset()
				if data.Invincible {
					ops.ColorM.ChangeHSV(0, 0, 1)
				}
				ops.GeoM.Translate( data.Ghosts[i].Position.PosX-float64(gwidth/2), (GridViewSize+data.GridOffsetY)-(data.Ghosts[i].Position.PosY+float64(gheight-(gheight/2))))
				if drawErr := view.DrawImage(ghostImg, ops); drawErr != nil {
					return nil, drawErr
				}
			}
			if state == GamePause {
				back, backErr := ebiten.NewImage(389, 130, ebiten.FilterDefault)
				if backErr != nil {
					return nil, backErr
				}
				if fillErr := back.Fill(color.Black); fillErr != nil {
					return nil, fillErr
				}
				text.Draw(back, "GAME PAUSED", fontface, 24, 65-(10), color.White)
				text.Draw(back, "PRESS SPACE", fontface, 24, 65+(10+31), color.White)
				ops.GeoM.Reset()
				ops.GeoM.Translate(370-(389/2), 512-(130/2))
				if drawErr := view.DrawImage(back, ops); drawErr != nil {
					return nil, drawErr
				}
			} else if state == GameOver {
				back, backErr := ebiten.NewImage(389, 130, ebiten.FilterDefault)
				if backErr != nil {
					return nil, backErr
				}
				if fillErr := back.Fill(color.Black); fillErr != nil {
					return nil, fillErr
				}

				text.Draw(back, "GAME OVER", fontface, 56, 65-(10), color.White)
				text.Draw(back, "PRESS SPACE", fontface, 24, 65+(10+31), color.White)

				ops.GeoM.Reset()
				ops.GeoM.Translate(370-(389/2), 512-(130/2))
				if drawErr := view.DrawImage(back, ops); drawErr != nil {
					return nil, drawErr
				}
			}
		}
		return view, nil
	}, nil
}

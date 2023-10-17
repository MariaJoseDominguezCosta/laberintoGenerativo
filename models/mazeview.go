// laberintogenerativo/models/mazeview.go
package models

import (
	"laberintogenerativo/resources"
	"laberintogenerativo/utils"
	"github.com/hajimehoshi/ebiten"
)

const MazeViewSize = 1536
const CellSize = 64

func MazeView(
	walls *resources.Walls,
) (func(state Mode, data *Data) (*ebiten.Image, error), error) {
	icWallSide, icWallSideErr := utils.ScaleSprite(walls.InActiveSide, 1.0, 1.0)
	if icWallSideErr != nil {
		return nil, icWallSideErr
	}
	icWallCorner, icWallCornerErr := utils.ScaleSprite(walls.InActiveCorner, 1.0, 1.0)
	if icWallCornerErr != nil {
		return nil, icWallCornerErr
	}
	mazeView, mazeViewErr := ebiten.NewImage(CellSize*Columns, MazeViewSize, ebiten.FilterDefault)
	if mazeViewErr != nil {
		return nil, mazeViewErr
	}
	var lastGrid [][Columns][4]rune
	return func(state Mode, data *Data) (*ebiten.Image, error) {
		if equal, copy := deepEqual(lastGrid, data.Grid); equal {
			return mazeView, nil
		} else {
			lastGrid = copy
		}
		if clearErr := mazeView.Clear(); clearErr != nil {
			return nil, clearErr
		}
		ops := &ebiten.DrawImageOptions{}
		for i := 0; i < len(data.Grid); i++ {
			for j := 0; j < len(data.Grid[i]); j++ {
				side := icWallSide
				corner := icWallCorner
				cellWalls := data.Grid[i][j]
				if cellWalls[0] == 'N' {
					ops.GeoM.Reset()
					ops.GeoM.Translate(float64(j*CellSize)+12, float64(MazeViewSize-((i*CellSize)+CellSize)))
					if drawErr := mazeView.DrawImage(side, ops); drawErr != nil {
						return nil, drawErr
					}
				}
				if cellWalls[1] == 'E' {
					ops.GeoM.Reset()
					ops.GeoM.Rotate(1.5708)
					ops.GeoM.Translate(float64(j*CellSize)+CellSize, float64(MazeViewSize-((i*CellSize)+52)))
					if drawErr := mazeView.DrawImage(side, ops); drawErr != nil {
						return nil, drawErr
					}
				}
				if cellWalls[2] == 'S' {
					ops.GeoM.Reset()
					ops.GeoM.Translate(float64(j*CellSize)+12, float64(MazeViewSize-((i*CellSize)+12)))
					if drawErr := mazeView.DrawImage(side, ops); drawErr != nil {
						return nil, drawErr
					}
				}
				if cellWalls[3] == 'W' {
					ops.GeoM.Reset()
					ops.GeoM.Rotate(1.5708)
					ops.GeoM.Translate(float64(j*CellSize)+12, float64(MazeViewSize-((i*CellSize)+52)))
					if drawErr := mazeView.DrawImage(side, ops); drawErr != nil {
						return nil, drawErr
					}
				}
				// Corners NE
				ops.GeoM.Reset()
				ops.GeoM.Translate(float64(j*CellSize)+52, float64(MazeViewSize-((i*CellSize)+CellSize)))
				if drawErr := mazeView.DrawImage(corner, ops); drawErr != nil {
					return nil, drawErr
				}
				// NW
				ops.GeoM.Reset()
				ops.GeoM.Translate(float64(j*CellSize), float64(MazeViewSize-((i*CellSize)+CellSize)))
				if drawErr := mazeView.DrawImage(corner, ops); drawErr != nil {
					return nil, drawErr
				}
				// SE
				ops.GeoM.Reset()
				ops.GeoM.Translate(float64(j*CellSize)+52, float64(MazeViewSize-((i*CellSize)+12)))
				if drawErr := mazeView.DrawImage(corner, ops); drawErr != nil {
					return nil, drawErr
				}
				// SW
				ops.GeoM.Reset()
				ops.GeoM.Translate(float64(j*CellSize), float64(MazeViewSize-((i*CellSize)+12)))
				if drawErr := mazeView.DrawImage(corner, ops); drawErr != nil {
					return nil, drawErr
				}
			}
		}
		return mazeView, nil
	}, nil
}

func deepEqual(previous, next [][Columns][4]rune) (bool, [][Columns][4]rune) {
	if len(previous) != len(next) {
		return false, copySlice(next)
	}
	for i := 0; i < len(previous); i++ {
		if previous[i] != next[i] {
			return false, copySlice(next)
		}
		for j := 0; j < Columns; j++ {
			if previous[i][j] != next[i][j] {
				return false, copySlice(next)
			}
		}
	}
	return true, next
}

func copySlice(src [][Columns][4]rune) [][Columns][4]rune {
	copy := make([][Columns][4]rune, len(src))
	for i := range src {
		copy[i] = src[i]
	}
	return copy
}

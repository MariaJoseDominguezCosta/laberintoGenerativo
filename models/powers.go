// laberintogenerativo/models/powers.go
package models

import "math"

const (
	Life powerType = iota
	Invincibility
)

func NewPower(x, y int, Kind powerType) Power {
	return Power{
		Position{
			CellX: x,
			CellY: y,
		},
		Kind,
	}
}
func (p *Data) PacmanTouchesPower(i int) bool {
	if p.Powers[i].Position.CellX == p.Player.Position.CellX && p.Powers[i].Position.CellY == p.Player.Position.CellY {
		posX := float64((p.Powers[i].Position.CellX * CellSize) + CellSize/2)
		posY := float64((p.Powers[i].Position.CellY * CellSize) + CellSize/2)
		if math.Abs(posX-p.Player.Position.PosX) < 20 && math.Abs(posY-(p.Player.Position.PosY+p.GridOffsetY)) < 20 {
			return true
		}
	}
	return false
}

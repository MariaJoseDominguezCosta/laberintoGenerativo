// laberintogenerativo/models/player.go
package models

import (
	"math"
	"math/rand"
)

type Player struct {
	Position
}

func NewPlayer() *Player {
	xcol := rand.Intn(Columns)
	return &Player{
		Position{CellX: xcol, CellY: 0,
			PosX: float64((xcol * CellSize) + (CellSize / 2)), PosY: CellSize / 2, Direction: North,
		},
	}
}

func CanMove(size float64, posX, posY float64, x, y int, walls [4]rune) bool {
	psx := posX - size
	psy := posY - size
	pex := posX + size
	pey := posY + size
	sx := x * CellSize
	sy := y * CellSize
	ex := sx + CellSize
	ey := sy + CellSize

	if walls[0] == 'N' {
		if pey > float64(ey-12) {
			return false
		}
	}
	if walls[1] == 'E' {
		if pex > float64(ex-12) {
			return false
		}
	}
	if walls[2] == 'S' {
		if psy < float64(sy+12) {
			return false
		}
	}
	if walls[3] == 'W' {
		if psx < float64(sx+12) {
			return false
		}
	}
	// NW corner
	if pey > float64(ey-12) && psx < float64(sx+12) {
		return false
	}
	// NE
	if pey > float64(ey-12) && pex > float64(ex-12) {
		return false
	}
	// SW
	if psy < float64(sy+12) && psx < float64(sx+12) {
		return false
	}
	// SE
	if psy < float64(sy+12) && pex > float64(ex-12) {
		return false
	}
	return true
}

// Move moves the player to a new position
func (p *Player) MovePlayer() {
	data := Data{}
	speed := 2.0
	switch data.Direction.Direction {
	case North, South:
		if p.PosX == float64((CellSize*p.CellX)+(CellSize/2)) {
			p.Direction = data.Direction.Direction
		}
	case East, West:
		if p.PosY+data.GridOffsetY == float64((CellSize*p.CellY)+(CellSize/2)) {
			p.Direction = data.Direction.Direction
		}
	}
	switch p.Direction {
	case North:
		if CanMove(
			20.0, p.PosX, p.PosY+data.GridOffsetY+speed, p.CellX, p.CellY, data.Grid[p.CellY][p.CellX],
		) {
			if p.PosY > OffsetY {
				data.GridOffsetY += speed
			} else {
				p.PosY += speed
			}
			if p.PosY+data.GridOffsetY+20 > float64((p.CellY*CellSize)+CellSize) {
				p.CellY += 1
			}
		}
	case South:
		if CanMove(20.0, p.PosX, p.PosY+data.GridOffsetY-speed, p.CellX, p.CellY, data.Grid[p.CellY][p.CellX]) {
			if p.PosY > OffsetY && data.GridOffsetY > 0 {
				data.GridOffsetY -= speed
			} else {
				p.PosY -= speed
			}
			if p.PosY+data.GridOffsetY-20 < float64(p.CellY*CellSize) {
				p.CellY -= 1
			}
		}
	case East:
		if CanMove(20.0, p.PosX+speed, p.PosY+data.GridOffsetY, p.CellX, p.CellY, data.Grid[p.CellY][p.CellX]) {
			p.PosX += speed
			if p.PosX+20 > float64((p.CellX*CellSize)+CellSize) {
				p.CellX += 1
			}
		}
	case West:
		if CanMove(20.0, p.PosX-speed, p.PosY+data.GridOffsetY, p.CellX, p.CellY, data.Grid[p.CellY][p.CellX]) {
			p.PosX -= speed
			if p.PosX-20 < float64(p.CellX*CellSize) {
				p.CellX -= 1
			}
		}
	}

}

func (p *Player) PacmanTouchesPower(i int) bool {
	data := Data{}
	if data.Powers[i].CellX == p.CellX && data.Powers[i].CellY == p.CellY {
		posX := float64((data.Powers[i].CellX * CellSize) + CellSize/2)
		posY := float64((data.Powers[i].CellY * CellSize) + CellSize/2)
		if math.Abs(posX-p.PosX) < 20 && math.Abs(posY-(p.PosY+data.GridOffsetY)) < 20 {
			return true
		}
	}
	return false
}

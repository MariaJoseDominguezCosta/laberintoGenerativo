// laberintogenerativo/models/ghosts.go
package models

import (
	"math"
	"math/rand"
)

type ghostType int

const (
	Ghost1 ghostType = iota
	Ghost2
	Ghost3
	Ghost4
)

// Image represents an image.
type Ghost struct {
	Position
	Kind ghostType
}

func NewGhost(x, y int, Kind ghostType, dir direction) Ghost {
	return Ghost{
		Position{
			CellX:     x,
			CellY:     y,
			PosX:      float64((x * CellSize) + CellSize/2),
			PosY:      float64((y * CellSize) + CellSize/2),
			Direction: dir,
		}, Kind,
	}
}

func (g Ghost) DirectionOfCell(cx, cy, nx, ny int) direction {
	if cx < nx {
		return East
	}
	if cx > nx {
		return West
	}
	if cy < ny {
		return North
	}
	if cy > ny {
		return South
	}

	if cx%2 == 0 {
		return West
	}
	return East
}

func (g Ghost) GetGhostDirection(i int) direction {
	data := Data{}
	pacX := float64((data.Player.CellX*CellSize) + (CellSize / 2))
	pacY := float64((data.Player.CellY*CellSize) + (CellSize / 2))

	ghost := data.Ghosts[i]

	if data.Player.CellX == ghost.CellX && data.Player.CellY == ghost.CellY {
		return ghost.Direction
	}

	x, y := ghost.CellX, ghost.CellY
	ghsX := float64((x*CellSize) + (CellSize / 2))
	ghsY := float64((y*CellSize) + (CellSize / 2))

	// since longest path can be m*n
	prevDist := float64((MazeViewSize/CellSize)*Columns) * CellSize
	if data.Invincible {
		prevDist = 0.0
	}

	for j := range rand.Perm(4) {
		if data.Grid[ghost.CellY][ghost.CellX][j] == '_' {
			nx, ny := 0, 0
			dist := 0.0
			switch j {
			case 0: // North
				// Added to prevent array overflow panic,
				// since last row in grid might have open North wall
				if y+1 < MazeViewSize/CellSize {
					dist = math.Sqrt(math.Pow(ghsX-pacX, 2) + math.Pow((ghsY+CellSize)-pacY, 2))
					nx, ny = ghost.CellX, ghost.CellY+1
				}
			case 1: // East
				dist = math.Sqrt(math.Pow((ghsX+CellSize)-pacX, 2) + math.Pow(ghsY-pacY, 2))
				nx, ny = ghost.CellX+1, ghost.CellY
			case 2: // South
				dist = math.Sqrt(math.Pow(ghsX-pacX, 2) + math.Pow((ghsY-CellSize)-pacY, 2))
				nx, ny = ghost.CellX, ghost.CellY-1
			case 3: // West
				dist = math.Sqrt(math.Pow((ghsX-CellSize)-pacX, 2) + math.Pow(ghsY-pacY, 2))
				nx, ny = ghost.CellX-1, ghost.CellY
			}
			if g.DirectionOfCell(ghost.CellX, ghost.CellY, nx, ny) !=
				GetOppositeDirection(ghost.Direction) {
				if data.Invincible {
					if dist > prevDist {
						x, y, prevDist = nx, ny, dist
					}
				} else {
					if dist < prevDist {
						x, y, prevDist = nx, ny, dist
					}
				}
			}
		}
	}

	return g.DirectionOfCell(ghost.CellX, ghost.CellY, x, y)
}

func (g Ghost) PacmanTouchesGhost(i int) bool {
	data := Data{}
	if data.Ghosts[i].CellX == data.Player.CellX && data.Ghosts[i].CellY == data.Player.CellY {
		posX := data.Ghosts[i].PosX
		posY := data.Ghosts[i].PosY
		if math.Abs(posX-data.Player.PosX) < 30 && math.Abs(posY-(data.Player.PosY+data.GridOffsetY)) < 30 {
			return true
		}
	}
	return false
}

func (g Ghost) MoveGhost(i int) {
	data := Data{}
	speed := 1.0
	ghost := data.Ghosts[i]

	if ghost.CellY >= MazeViewSize/CellSize {
		return
	}

	if IsIntersection(data.Grid[ghost.CellY][ghost.CellX]) {
		if ghost.PosX == float64((CellSize*ghost.CellX)+(CellSize/2)) && ghost.PosY == float64((CellSize*ghost.CellY)+(CellSize/2)) {
			data.Ghosts[i].Direction = g.GetGhostDirection(i)
		}
	} else if IsBlocked(data.Grid[ghost.CellY][ghost.CellX], ghost.Direction) || IsDeadend(data.Grid[ghost.CellY][ghost.CellX]) {
		if ghost.PosX == float64((CellSize*ghost.CellX)+(CellSize/2)) && ghost.PosY == float64((CellSize*ghost.CellY)+(CellSize/2)) {
			data.Ghosts[i].Direction = GetExit(data.Grid[ghost.CellY][ghost.CellX])
		}
	}

	switch data.Ghosts[i].Direction {
	case North:
		if CanMove(
			20.0,
			data.Ghosts[i].PosX,
			data.Ghosts[i].PosY+speed,
			data.Ghosts[i].CellX,
			data.Ghosts[i].CellY,
			data.Grid[data.Ghosts[i].CellY][data.Ghosts[i].CellX],
		) {
			data.Ghosts[i].PosY += speed
			if data.Ghosts[i].PosY+20 > float64((data.Ghosts[i].CellY*CellSize)+CellSize) {
				data.Ghosts[i].CellY += 1
			}
		}
	case South:
		if CanMove(
			20.0,
			data.Ghosts[i].PosX,
			data.Ghosts[i].PosY-speed,
			data.Ghosts[i].CellX,
			data.Ghosts[i].CellY,
			data.Grid[data.Ghosts[i].CellY][data.Ghosts[i].CellX],
		) {
			data.Ghosts[i].PosY -= speed
			if data.Ghosts[i].PosY-20 < float64((data.Ghosts[i].CellY*CellSize)) {
				data.Ghosts[i].CellY -= 1
			}
		}
	case East:
		if CanMove(
			20.0,
			data.Ghosts[i].PosX+speed,
			data.Ghosts[i].PosY,
			data.Ghosts[i].CellX,
			data.Ghosts[i].CellY,
			data.Grid[ghost.CellY][ghost.CellX],
		) {
			data.Ghosts[i].PosX += speed
			if data.Ghosts[i].PosX+20 > float64((data.Ghosts[i].CellX*CellSize)+CellSize) {
				data.Ghosts[i].CellX += 1
			}
		}
	case West:
		if CanMove(
			20.0,
			data.Ghosts[i].PosX-speed,
			data.Ghosts[i].PosY,
			data.Ghosts[i].CellX,
			data.Ghosts[i].CellY,
			data.Grid[data.Ghosts[i].CellY][data.Ghosts[i].CellX],
		) {
			data.Ghosts[i].PosX -= speed
			if data.Ghosts[i].PosX-20 < float64(data.Ghosts[i].CellX*CellSize) {
				data.Ghosts[i].CellX -= 1
			}
		}
	}
}

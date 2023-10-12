// laberintogenerativo/models/ghosts.go
package models

import (
	"math"
	"math/rand"
)

func (g *Data) GetGhostDirection(i int) Direction {
	pacX := float64((g.Player.Position.CellX*CellSize)+(CellSize/2))
	pacY := float64((g.Player.Position.CellY*CellSize)+(CellSize/2))
	ghost := g.Ghosts[i]
	if ((g.Player.Position.CellX) == (ghost.Position.CellX) && (g.Player.Position.CellY) == (ghost.Position.CellY)){
		return (ghost.Position.Direction)
	}
	x, y := (ghost.Position.CellX), (ghost.Position.CellY)
	ghsX := float64((x*CellSize)+(CellSize/2))
	ghsY := float64((y*CellSize)+(CellSize/2))
	// since longest path can be m*n
	prevDist := float64((MazeViewSize/CellSize)*Columns)*CellSize
	if g.Invincible {
		prevDist = 0.0
	}
	for j := range rand.Perm(4) {
		if g.Grid[ghost.Position.CellY][ghost.Position.CellX][j] == '_' {
			nx, ny := 0, 0
			dist := 0.0
			switch j {
			case 0: // North
				if y+1 < MazeViewSize/CellSize {
					dist = math.Sqrt(math.Pow(ghsX-pacX, 2)+math.Pow((ghsY+CellSize)-pacY, 2))
					nx, ny = ghost.CellX, ghost.CellY+1
				}
			case 1: // East
				dist = math.Sqrt(math.Pow((ghsX+CellSize)-pacX, 2)+math.Pow(ghsY-pacY, 2))
				nx, ny = ghost.Position.CellX+1, ghost.Position.CellY
			case 2: // South
				dist = math.Sqrt(math.Pow(ghsX-pacX, 2)+math.Pow((ghsY-CellSize)-pacY, 2))
				nx, ny = ghost.Position.CellX, ghost.Position.CellY-1
			case 3: // West
				dist = math.Sqrt(math.Pow((ghsX-CellSize)-pacX, 2)+math.Pow(ghsY-pacY, 2))
				nx, ny = ghost.Position.CellX-1, ghost.Position.CellY
			}
			if g.DirectionOfCell(ghost.Position.CellX, ghost.Position.CellY, nx, ny) != GetOppositeDirection(ghost.Position.Direction) {
				if g.Invincible {
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
	return g.DirectionOfCell(ghost.Position.CellX, ghost.Position.CellY, x, y)
}
func (g *Data) PacmanTouchesGhost(i int) bool {
	if (((g.Ghosts[i].Position.CellX) == (g.Player.Position.CellX) )&& ((g.Ghosts[i].Position.CellY) == (g.Player.Position.CellY))) {
		posX := (g.Ghosts[i].Position.PosX)
		posY := (g.Ghosts[i].Position.PosY)
		if math.Abs(posX-(g.Player.Position.PosX)) < 30 && math.Abs(posY-((g.Player.Position.PosY)+(g.GridOffsetY))) < 30 {
			return true
		}
	}
	return false
}

func (g *Data) MoveGhost(i int) {
	speed := 1.0
	ghost := g.Ghosts[i]
	if ((ghost.Position.CellY) >= (MazeViewSize/CellSize)) {
		return
	}

	if IsIntersection(g.Grid[ghost.Position.CellY][ghost.Position.CellX]) {
		if ((ghost.Position.PosX) == float64((CellSize*(ghost.Position.CellX))+(CellSize/2)) && (ghost.Position.PosY) == float64((CellSize*(ghost.Position.CellY))+(CellSize/2))) {
			(g.Ghosts[i].Position.Direction) = g.GetGhostDirection(i)
		}
	} else if IsBlocked(g.Grid[ghost.Position.CellY][ghost.Position.CellX], ghost.Position.Direction) || IsDeadend(g.Grid[ghost.Position.CellY][ghost.Position.CellX]) {
		if((ghost.Position.PosX) == float64((CellSize*(ghost.Position.CellX))+(CellSize/2)) && (ghost.Position.PosY) == float64((CellSize*(ghost.Position.CellY))+(CellSize/2))) {
			(g.Ghosts[i].Position.Direction) = GetExit(g.Grid[ghost.Position.CellY][ghost.Position.CellX])
		}
	}

	switch (g.Ghosts[i].Position.Direction ){
	case North:
		if CanMove(
			20.0,
			(g.Ghosts[i].Position.PosX),
			(g.Ghosts[i].Position.PosY)+speed,
			(g.Ghosts[i].Position.CellX),
			(g.Ghosts[i].Position.CellY),
			g.Grid[g.Ghosts[i].Position.CellY][g.Ghosts[i].Position.CellX],
		) {
			(g.Ghosts[i].Position.PosY) += speed
			if ((g.Ghosts[i].Position.PosY+20) > float64(((g.Ghosts[i].Position.CellY)*CellSize)+CellSize)) {
				(g.Ghosts[i].Position.CellY) += 1
			}
		}
	case South:
		if CanMove(
			20.0,
			(g.Ghosts[i].Position.PosX),
			(g.Ghosts[i].Position.PosY)-speed,
			(g.Ghosts[i].Position.CellX),
			(g.Ghosts[i].Position.CellY),
			g.Grid[g.Ghosts[i].Position.CellY][g.Ghosts[i].Position.CellX],
		) {
			(g.Ghosts[i].Position.PosY) -= speed
			if ((g.Ghosts[i].Position.PosY)-20 < float64(((g.Ghosts[i].Position.CellY)*CellSize))) {
				(g.Ghosts[i].Position.CellY) -= 1
			}
		}
	case East:
		if CanMove(
			20.0,
			(g.Ghosts[i].Position.PosX)+speed,
			(g.Ghosts[i].Position.PosY),
			(g.Ghosts[i].Position.CellX),
			(g.Ghosts[i].Position.CellY),
			g.Grid[ghost.Position.CellY][ghost.Position.CellX],
		) {
			(g.Ghosts[i].Position.PosX) += speed
			if( (g.Ghosts[i].Position.PosX)+20 > float64(((g.Ghosts[i].Position.CellX)*CellSize)+CellSize)) {
				(g.Ghosts[i].Position.CellX) += 1
			}
		}
	case West:
		if CanMove(
			20.0,
			(g.Ghosts[i].Position.PosX)-speed,
			g.Ghosts[i].Position.PosY,
			g.Ghosts[i].Position.CellX,
			g.Ghosts[i].Position.CellY,
			g.Grid[g.Ghosts[i].Position.CellY][g.Ghosts[i].Position.CellX],
		) {
			(g.Ghosts[i].Position.PosX) -= speed
			if ((g.Ghosts[i].Position.PosX)-20 < float64((g.Ghosts[i].Position.CellX)*CellSize)) {
				(g.Ghosts[i].Position.CellX) -= 1
			}
		}
	}
}

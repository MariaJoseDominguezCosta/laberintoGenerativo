// laberintogenerativo/models/data.go
package models

type Direction int
type powerType int
type ghostType int
type Mode int
type Data struct {
	Grid        [][Columns][4]rune
	Active      [][Columns]bool
	Lifes       int
	Score       int
	Player      Player
	Ghosts      []Ghost
	Powers      []Power
	GridOffsetY float64
	Invincible  bool
	Direction   Direction
	Mode        Mode
}
type Power struct {
	Position
	Kind powerType
}
type Player struct {
	Position
}
type Position struct {
	CellX, CellY int
	PosX, PosY   float64
	Direction    Direction
}
type Ghost struct {
	Position
	Kind ghostType
}

const (
	Ghost1 ghostType = iota
	Ghost2
	Ghost3
	Ghost4
)
const (
	GameLoading Mode = iota
	GameStart
	GamePause
	GameOver
	GridOffsetY = (CellSize * 10)
)
const (
	North Direction = iota
	East
	South
	West
)

func NewData() *Data {
	return &Data{
		Lifes: 5,
		Score: 1,
	}
}
func NewGhost(x, y int, Kind ghostType, dir Direction) Ghost {
	posX := float64((x * CellSize) + CellSize/2)
	posY := float64((y * CellSize) + CellSize/2)
	return Ghost{
		Position{
			CellX:     x,
			CellY:     y,
			PosX:      posX,
			PosY:      posY,
			Direction: dir,
		}, Kind,
	}
}
func (g *Data) DirectionOfCell(cx, cy, nx, ny int) Direction {
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
func GetOppositeDirection(dir Direction) Direction {
	switch dir {
	case North:
		return South
	case East:
		return West
	case South:
		return North
	default:
		return East
	}
}
func NumOfWalls(walls [4]rune) int {
	count := 0
	if walls[0] == 'N' {
		count += 1
	}
	if walls[1] == 'E' {
		count += 1
	}
	if walls[2] == 'S' {
		count += 1
	}
	if walls[3] == 'W' {
		count += 1
	}
	return count
}
func IsIntersection(walls [4]rune) bool {
	count := NumOfWalls(walls)
	if count >= 3 {
		return false
	} else if count == 2 {
		// covers the case of corridor
		if walls[0] == walls[2] || walls[1] == walls[3] {
			return false
		}
	}
	return true
}
func (g *Data) Keyboard() {
	if g == nil {
		return
	}
	walls := g.Grid[g.Player.Position.CellY][g.Player.Position.CellX]
	if upKeyPressed() && walls[0] == '_' {
		g.Direction = North
	}
	if downKeyPressed() && walls[2] == '_' {
		g.Direction = South
	}
	if leftKeyPressed() && walls[3] == '_' {
		g.Direction = West
	}
	if rightKeyPressed() && walls[1] == '_' {
		g.Direction = East
	}
}
func IsDeadend(walls [4]rune) bool {
	return NumOfWalls(walls) >= 3
}
func GetExit(walls [4]rune) Direction {
	for i := 0; i < 4; i++ {
		if walls[i] == '_' {
			switch i {
			case 0:
				return North
			case 1:
				return East
			case 2:
				return South
			case 3:
				return West
			}
		}
	}
	return North
}
func IsBlocked(walls [4]rune, dir Direction) bool {
	switch dir {
	case North:
		return walls[0] != '_'
	case East:
		return walls[1] != '_'
	case South:
		return walls[2] != '_'
	case West:
		return walls[3] != '_'
	}
	return false
}
func CanMove(size float64, posX, posY float64, x, y int, walls [4]rune) bool {
	switch {
	case walls[0] == 'N' && posY-size > float64(y*CellSize+12):
		return false
	case walls[1] == 'E' && posX+size > float64(x*CellSize+CellSize-12):
		return false
	case walls[2] == 'S' && posY+size < float64(y*CellSize+CellSize-12):
		return false
	case walls[3] == 'W' && posX-size < float64(x*CellSize+12):
		return false
	case posY-size > float64(y*CellSize+12) && posX-size < float64(x*CellSize+12):
		return false
	case posY-size > float64(y*CellSize+12) && posX+size > float64(x*CellSize+CellSize-12):
		return false
	case posY+size < float64(y*CellSize+CellSize-12) && posX-size < float64(x*CellSize+12):
		return false
	case posY+size < float64(y*CellSize+CellSize-12) && posX+size > float64(x*CellSize+CellSize-12):
		return false
	}
	return true
}

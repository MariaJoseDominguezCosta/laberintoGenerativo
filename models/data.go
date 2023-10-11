// laberintogenerativo/models/data.go

package models

type direction int

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
	Direction   Position
	Mode        Mode
}

const MaxLifes = 7
const (
	GameLoading Mode = iota
	GameStart
	GamePause
	GameOver

	OffsetY = CellSize * 10
)

const (
	North direction = iota
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

type Position struct {
	CellX, CellY int
	PosX, PosY   float64
	Direction    direction
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

func GetOppositeDirection(dir direction) direction {
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

func IsBlocked(walls [4]rune, dir direction) bool {
	switch dir {
	case North:
		if walls[0] != '_' {
			return true
		}
	case East:
		if walls[1] != '_' {
			return true
		}
	case South:
		if walls[2] != '_' {
			return true
		}
	case West:
		if walls[3] != '_' {
			return true
		}
	}
	return false
}

func GetExit(walls [4]rune) direction {
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

func IsDeadend(walls [4]rune) bool {
	return NumOfWalls(walls) >= 3
}
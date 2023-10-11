// laberintogenerativo/models/powers.go
package models

type powerType int

type Power struct {
	Position
	Kind powerType
}

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
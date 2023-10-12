// laberintogenerativo/models/powers.go
package models

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

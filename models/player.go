// laberintogenerativo/models/player.go
package models
// Move moves the player to a new position
func (p *Data) MovePlayer() {
	speed := 2.0
	xcell := p.Player.Position.CellX
	ycell := p.Player.Position.CellY
	switch p.Direction {
	case North, South:
		if p.Player.Position.PosX == float64((CellSize*xcell)+(CellSize/2)) {
			p.Player.Position.Direction = p.Direction
		}
	case East, West:
		if p.Player.Position.PosY+p.GridOffsetY == float64((CellSize*ycell)+(CellSize/2)) {
			p.Player.Position.Direction = p.Direction
		}
	}
	switch p.Player.Position.Direction {
	case North:
		if CanMove(
			20.0,
			p.Player.Position.PosX,
			p.Player.Position.PosY+p.GridOffsetY+speed,
			p.Player.Position.CellX,
			p.Player.Position.CellY,
			p.Grid[ycell][xcell],
		) {
			if p.Player.Position.PosY > GridOffsetY {
				p.GridOffsetY += speed
			} else {
				p.Player.Position.PosY += speed
			}
			if p.Player.Position.PosY+p.GridOffsetY+20 > float64((ycell*CellSize)+CellSize) {
				p.Player.Position.CellY += 1
			}
		}
	case South:
		if CanMove(
			20.0,
			p.Player.Position.PosX,
			p.Player.Position.PosY+p.GridOffsetY-speed,
			p.Player.Position.CellX,
			p.Player.Position.CellY,
			p.Grid[ycell][xcell],
		) {
			if p.Player.Position.PosY > GridOffsetY && p.GridOffsetY > 0 {
				p.GridOffsetY -= speed
			} else {
				p.Player.Position.PosY -= speed
			}
			if p.Player.Position.PosY+p.GridOffsetY-20 < float64(ycell*CellSize) {
				p.Player.Position.CellY -= 1
			}
		}
	case East:
		if CanMove(
			20.0,
			p.Player.Position.PosX+speed,
			p.Player.Position.PosY+p.GridOffsetY,
			p.Player.Position.CellX,
			p.Player.Position.CellY,
			p.Grid[ycell][xcell],
		) {
			p.Player.Position.PosX += speed
			if p.Player.Position.PosX+20 > float64((xcell*CellSize)+CellSize) {
				p.Player.Position.CellX += 1
			}
		}
	case West:
		if CanMove(
			20.0,
			p.Player.Position.PosX-speed,
			p.Player.Position.PosY+p.GridOffsetY,
			p.Player.Position.CellX,
			p.Player.Position.CellY,
			p.Grid[ycell][xcell],
		) {
			p.Player.Position.PosX -= speed
			if p.Player.Position.PosX-20 < float64(xcell*CellSize) {
				p.Player.Position.CellX -= 1
			}
		}
	}
}

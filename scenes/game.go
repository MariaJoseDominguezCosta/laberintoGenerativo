// laberintogenerativo/scenes/game.go
package scenes

import (
	"laberintogenerativo/models"
	"laberintogenerativo/resources"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	state       models.Mode
	rand        *rand.Rand
	maze        *models.Maze
	data        *models.Data
	ghost       *models.Ghost
	gridView    func(models.Mode, *models.Data) (*ebiten.Image, error)
	direction   *models.Position
	powerTicker *time.Ticker
}

// Layout implements ebiten.Game.
func (*Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenWidth, screenHeight
}
func NewGame() (*Game, error) {

	lAssets, assetsErr := resources.LoadAssets()
	if assetsErr != nil {
		return nil, assetsErr
	}

	mazeView, mazeViewErr := models.MazeView(lAssets.Walls)
	if mazeViewErr != nil {
		return nil, mazeViewErr
	}

	gridView, gridViewErr := models.GridView(lAssets.Characters, lAssets.Powers,
		lAssets.ArcadeFont, mazeView)
	if gridViewErr != nil {
		return nil, gridViewErr
	}

	return &Game{
		rand:     rand.New(rand.NewSource(time.Now().UnixNano())),
		state:    models.GameLoading,
		gridView: gridView,
	}, nil
}

// Update implements ebiten.Game.
func (g *Game) Update(screen *ebiten.Image) error {
	switch g.state {
	case models.GameLoading:
		if models.SpaceReleased() {
			xcol := g.rand.Intn(models.Columns)
			numOfRows := models.MazeViewSize / models.CellSize
			g.data = models.NewData()
			g.maze = models.NewPopulatedMaze(32, g.rand)
			g.data.Grid = g.maze.Get(0, numOfRows)
			g.data.Active = make([][models.Columns]bool, numOfRows)
			g.data.Player = *models.NewPlayer()
			g.direction = &g.data.Player.Position
			g.data.Active[0][xcol] = true

			powers := make([]models.Power, 0)
			for i := 0; i < numOfRows; i += 4 {
				cellX := g.rand.Intn(models.Columns)
				cellY := g.rand.Intn(4) + i
				kind := models.Invincibility
				if (cellY-i)%2 == 0 {
					kind = models.Life
				}
				powers = append(powers, models.NewPower(cellX, cellY, kind))
			}
			g.data.Powers = powers

			ghosts := make([]models.Ghost, 0)
			for i := 0; i < numOfRows; i += 2 {
				cellX := g.rand.Intn(models.Columns/2) + models.Columns/2
				if i%4 == 0 {
					cellX = g.rand.Intn(models.Columns / 2)
				}
				cellY := g.rand.Intn(2) + i
				kind := models.Ghost1
				if (cellY-i)%4 == 0 {
					kind = models.Ghost4
				} else if (cellY-i)%3 == 0 {
					kind = models.Ghost3
				} else if (cellY-i)%2 == 0 {
					kind = models.Ghost2
				}
				ghosts = append(ghosts, models.NewGhost(cellX, cellY, kind, models.GetExit(
					g.data.Grid[cellY][cellX])))
			}
			g.data.Ghosts = ghosts

			g.state = models.GameStart
		} else {
			g.data = nil
			g.maze = nil
		}
	case models.GameStart:
		if models.SpaceReleased() {
			g.state = models.GamePause
		} else if g.data.Lifes < 1 {
			g.state = models.GameOver
		} else {
			numOfRows := models.MazeViewSize / models.CellSize
			if g.data.Player.CellY == len(g.data.Grid)-8 {
				g.maze.Compact(4)
				if (g.maze.Rows() - numOfRows) < 4 {
					g.maze.GrowBy(16)
				}

				g.data.Grid = g.maze.Get(0, numOfRows)
				// shift active grid by 4
				for i := 4; i <= len(g.data.Active); i++ {
					for j := 0; j < models.Columns; j++ {
						if i <= g.data.Player.CellY {
							g.data.Active[i-4][j] = g.data.Active[i][j]
						} else {
							g.data.Active[i-4][j] = false
						}
					}
				}

				g.data.Player.CellY -= 4
				g.data.GridOffsetY -= models.CellSize * 4

				for i := 0; i < len(g.data.Powers); i++ {
					g.data.Powers[i].CellY -= 4
					if g.data.Powers[i].CellY < 0 {
						cellX := g.rand.Intn(models.Columns)
						cellY := g.rand.Intn(4) + (numOfRows - 4)
						g.data.Powers[i] = models.NewPower(cellX, cellY, g.data.Powers[i].Kind)
					}
				}
				for i := 0; i < len(g.data.Ghosts); i++ {
					g.data.Ghosts[i].CellY -= 4
					g.data.Ghosts[i].PosY -= models.CellSize * 4
					if g.data.Ghosts[i].CellY < 0 {
						cellX := g.rand.Intn(models.Columns)
						cellY := g.rand.Intn(4) + (numOfRows - 4)
						g.data.Ghosts[i] = models.NewGhost(
							cellX, cellY,
							g.data.Ghosts[i].Kind,
							models.GetExit(g.data.Grid[cellY][cellX]))
					}
				}
			}

			g.Keyboard()
			g.data.Player.MovePlayer()

			if !g.data.Active[g.data.Player.CellY][g.data.Player.CellX] {
				if math.Abs(float64(
					(g.data.Player.CellX * models.CellSize)+(models.CellSize/2),
				)-(g.data.Player.PosX)) < 20 && math.Abs(float64(
						(g.data.Player.CellY * models.CellSize)+(models.CellSize/2),
					)-(g.data.Player.PosY + g.data.GridOffsetY)) < 20 {
					g.data.Active[g.data.Player.CellY][g.data.Player.CellX] = true
					g.data.Score += 1
				}
			}

			// check powers
			for i := 0; i < len(g.data.Powers); i++ {
				cellX := g.rand.Intn(models.Columns)
				cellY := g.rand.Intn(4) + (((g.data.Powers[i].CellY / 4) * 4) + numOfRows)
				if g.data.Player.PacmanTouchesPower(i) {
					switch g.data.Powers[i].Kind {
					case models.Life:
						if g.data.Lifes < models.MaxLifes {
							g.data.Lifes += 1
							g.data.Powers[i] = models.NewPower(cellX, cellY, g.data.Powers[i].Kind)
						}
					case models.Invincibility:
						if !g.data.Invincible {
							g.data.Invincible = true
						}
						g.StartCountdown(10)
						g.data.Powers[i] = models.NewPower(cellX, cellY, g.data.Powers[i].Kind)
					}
				}
			}
			// check ghosts
			for i := 0; i < len(g.data.Ghosts); i++ {
				if g.ghost.PacmanTouchesGhost(i) {
					if !g.data.Invincible {
						g.data.Lifes -= 1
					} else {
						g.data.Score += 200
					}
					cellX := g.rand.Intn(models.Columns)
					cellY := g.rand.Intn(4) + (((g.data.Ghosts[i].CellY / 4) * 4) + numOfRows)
					g.data.Ghosts[i] = models.NewGhost(
						cellX, cellY, g.data.Ghosts[i].Kind, models.North)
				}
				g.ghost.MoveGhost(i)
			}
		}
	case models.GamePause:
		if models.SpaceReleased() {
			g.state = models.GameStart
		}
	case models.GameOver:
		if models.SpaceReleased() {
			g.state = models.GameLoading

		}
	default:
		// reset state to GameLoading
		// dont return error
		g.state = models.GameLoading
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	gview, gviewErr := g.gridView(g.state, g.data)
	if gviewErr != nil {
		return gviewErr
	}

	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Reset()
	ops.GeoM.Translate(38, 162)
	if drawErr := screen.DrawImage(gview, ops); drawErr != nil {
		return drawErr
	}

	return nil
}

func (g *Game) Run() error {
	return ebiten.Run(func(screen *ebiten.Image) error {
		return g.Update(screen)
	}, 712, 1220, 0.5, "PACMAN") // scale is kept to 0.5, for good rendering in retina.
}
func (g *Game) Keyboard() {
	if g.data != nil {
		walls := g.data.Grid[g.data.Player.CellY][g.data.Player.CellX]
		if models.UpKeyPressed() {
			if walls[0] == '_' {
				g.data.Direction.Direction = models.North
			}
		}
		if models.DownKeyPressed() {
			if walls[2] == '_' {
				g.data.Direction.Direction = models.South
			}
		}
		if models.LeftKeyPressed() {
			if walls[3] == '_' {
				g.data.Direction.Direction = models.West
			}
		}
		if models.RightKeyPressed() {
			if walls[1] == '_' {
				g.data.Direction.Direction = models.East
			}
		}
	}
}

func (g *Game) StartCountdown(duration int) {
	if g.powerTicker != nil {
		g.powerTicker.Stop()
	}
	g.powerTicker = time.NewTicker(time.Duration(duration) * time.Second)
	go func() {
		select {
		case <-g.powerTicker.C:
			g.data.Invincible = false
		}
	}()
}
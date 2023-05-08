package antique

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/steambap/antique/ui"
)

type Game struct {
	Day       float64
	GameSpeed float64
	Galaxy    *Galaxy
	UI        *ui.UI
}

var TheGame *Game

func NewGame() {
	TheGame = &Game{
		Day:       0,
		GameSpeed: 2,
		Galaxy:    &Galaxy{},
	}

	TheGame.UI = ui.New()
	TheGame.Galaxy.Init()
}

func (g *Game) Update() error {
	g.Day += 0.01 * g.GameSpeed
	g.Galaxy.Update()
	g.UI.StatusPanel.DayLabel.UpdateText(fmt.Sprintf("Day: %06d", int(g.Day)))
	g.UI.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Galaxy.Draw(screen)
	g.UI.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	g.UI.Layout(outsideWidth, outsideHeight)
	return 640, 480
}

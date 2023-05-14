package antique

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/steambap/antique/generator"
)

type Galaxy struct {
	Points []gg.Point
	camX   float64
	camY   float64
	img    *ebiten.Image
}

func (g *Galaxy) Init() {
	g.Points = generator.NewGalaxy()
	dc := gg.NewContext(500, 500)
	for _, point := range g.Points {
		dc.Push()
		dc.DrawPoint(point.X*500, point.Y*500, 1)
		dc.SetHexColor("#ff7f22")
		dc.Fill()
		dc.Pop()
	}

	g.img = ebiten.NewImageFromImage(dc.Image())
}

func (g *Galaxy) Update() {
	// Pan camera via keyboard.
	pan := 7.0 / 1.0
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.camX += pan
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.camX -= pan
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.camY -= pan
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.camY += pan
	}
}

func (g *Galaxy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.camX, g.camY)
	screen.DrawImage(g.img, op)
}

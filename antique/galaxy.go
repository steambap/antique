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
}

func (g *Galaxy) Init() {
	g.Points = generator.NewGalaxy(0.5, 4)

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
	dc := gg.NewContext(640, 480)
	for _, point := range g.Points {
		dc.Push()
		dc.DrawPoint(g.camX+point.X*500, g.camY+point.Y*500, 1)
		dc.SetHexColor("#ff7f22")
		dc.Fill()
		dc.Pop()
	}

	img := ebiten.NewImageFromImage(dc.Image())
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(img, op)
}

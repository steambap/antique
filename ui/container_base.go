package ui

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/steambap/antique/util"
)

type ContainerBase struct {
	img     *ebiten.Image
	widgets []Widget
	x, y    int
	width   int
	height  int
}

func (cb *ContainerBase) createImg(color color.Color) *ebiten.Image {
	if cb.width == 0 || cb.height == 0 {
		return nil
	}

	dc := gg.NewContext(cb.width, cb.height)
	dc.SetColor(color)
	dc.DrawRectangle(0, 0, float64(cb.width), float64(cb.height))
	dc.Fill()

	return ebiten.NewImageFromImage(dc.Image())
}

func (cb *ContainerBase) Position() (int, int) {
	return cb.x, cb.y
}

func (cb *ContainerBase) Size() (int, int) {
	return cb.width, cb.height
}

func (cb *ContainerBase) Rect() (int, int, int, int) {
	x0 := cb.x
	y0 := cb.y
	x1 := x0 + cb.width
	y1 := y0 + cb.height
	return x0, y0, x1, y1
}

func (cb *ContainerBase) Widgets() []Widget {
	return cb.widgets
}

func (cb *ContainerBase) FindWidgetAt(x, y int) Widget {
	for _, w := range cb.widgets {
		if util.InRect(x, y, cb.Rect) {
			return w
		}
	}

	return nil
}

func (cb *ContainerBase) Draw(screen *ebiten.Image) {
	if cb.img == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(cb.x), float64(cb.y))
	screen.DrawImage(cb.img, op)

	for _, w := range cb.widgets {
		w.Draw(screen)
	}
}

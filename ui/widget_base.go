package ui

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

// Widget base is a button that display a single icon
type WidgetBase struct {
	parent   Container
	id       string
	img      *ebiten.Image
	align    gg.Align
	disabled bool
	x, y     int
	width    int
	height   int
}

func (wb *WidgetBase) Parent() Container {
	return wb.parent
}

func (wb *WidgetBase) ID() string {
	return wb.id
}

func (wb *WidgetBase) Size() (int, int) {
	return wb.width, wb.height
}

func (wb *WidgetBase) Position() (int, int) {
	return wb.x, wb.y
}

func (wb *WidgetBase) Rect() (int, int, int, int) {
	x0 := wb.x
	y0 := wb.y
	x1 := x0 + wb.width
	y1 := y0 + wb.height
	return x0, y0, x1, y1
}

func (wb *WidgetBase) OffsetRect() (int, int, int, int) {
	px, py := wb.parent.Position()
	x0 := wb.x + px
	y0 := wb.y + py
	x1 := x0 + wb.width
	y1 := y0 + wb.height
	return x0, y0, x1, y1
}

func (wb *WidgetBase) SetPosition(x, y int) {
	wb.x = x
	wb.y = y
}

func (wb *WidgetBase) Align() gg.Align {
	return wb.align
}

func (wb *WidgetBase) Disabled() bool {
	return wb.disabled
}

func (wb *WidgetBase) Activate() {
	panic("not implemented") // TODO: Implement
}

func (wb *WidgetBase) Deactivate() {
	panic("not implemented") // TODO: Implement
}

func (wb *WidgetBase) Tapped() {
	panic("not implemented") // TODO: Implement
}

func (wb *WidgetBase) Update() {}

func (wb *WidgetBase) Draw(screen *ebiten.Image) {
	if wb.img == nil {
		return
	}

	// Don't draw a widget unless it is fully contained within it's parent
	_, pt, _, pb := wb.parent.Rect()
	widgetLeft, widgetTop, _, widgetBottom := wb.OffsetRect()
	if widgetBottom > pb || widgetBottom-wb.height < pt {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(widgetLeft), float64(widgetTop))
	if wb.disabled {
		op.ColorScale.Scale(1, 1, 1, 0.25)
	}

	screen.DrawImage(wb.img, op)
}

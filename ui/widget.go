package ui

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

// IWidget is an interface for widget
type Widget interface {
	Parent() Container
	ID() string
	Size() (int, int)
	Position() (int, int)
	Rect() (int, int, int, int)
	OffsetRect() (int, int, int, int)
	SetPosition(int, int)
	Align() gg.Align
	Disabled() bool
	Activate()
	Deactivate()
	Tapped()
	Update()
	Draw(*ebiten.Image)
}

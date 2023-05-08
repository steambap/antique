package ui

import "github.com/hajimehoshi/ebiten/v2"

type UI struct {
	StatusPanel *StatusPanel
	containers  []Container
}

func New() *UI {
	ui := &UI{}

	sp := NewStatusPanel()
	ui.StatusPanel = sp
	ui.containers = []Container{sp}

	return ui
}

func (u *UI) Layout(outsideWidth, outsideHeight int) (int, int) {
	for _, c := range u.containers {
		c.Layout(outsideWidth, outsideHeight)
	}

	return outsideWidth, outsideHeight
}

func (u *UI) Update() {
	for _, c := range u.containers {
		c.Update()
	}
}

func (u *UI) Draw(screen *ebiten.Image) {
	for _, c := range u.containers {
		c.Draw(screen)
	}
}

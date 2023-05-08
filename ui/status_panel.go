package ui

import (
	"github.com/fogleman/gg"
	"github.com/steambap/antique/typeface"
)

type StatusPanel struct {
	DayLabel *Label
	PanelBase
}

func NewStatusPanel() *StatusPanel {
	sp := &StatusPanel{
		PanelBase: PanelBase{
			ContainerBase{
				x: 0, y: 0,
				width:  0,
				height: 48,
			},
		},
	}

	sp.DayLabel = NewLabel(sp, "Date", gg.AlignRight, "Day: {}", typeface.GoRegular14, "")
	sp.widgets = []Widget{
		sp.DayLabel,
	}

	return sp
}

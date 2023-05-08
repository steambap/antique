package ui

import "github.com/fogleman/gg"

const PADDING_SIZE = 24

type PanelBase struct {
	ContainerBase
}

// Layout widgets that belong to this container
// by setting the x,y of each relative to their parent
func (pb *PanelBase) LayoutWidgets() {
	nextLeft := PADDING_SIZE
	nextRight := pb.width - PADDING_SIZE
	for _, w := range pb.widgets {
		widgetWidth, widgetHeight := w.Size()
		parentHeight := pb.height
		y := parentHeight/2 - widgetHeight/2
		switch w.Align() {
		case gg.AlignLeft:
			w.SetPosition(nextLeft, y)
			nextLeft += PADDING_SIZE
		case gg.AlignCenter:
			w.SetPosition(pb.width/2-widgetWidth/2, y)
		case gg.AlignRight:
			w.SetPosition(nextRight-widgetWidth, y)
			nextRight -= widgetWidth + PADDING_SIZE
		}
	}
}

func (pb *PanelBase) StartDrag() {}

func (pb *PanelBase) DragBy(dx, dy int) {}

func (pb *PanelBase) StopDrag() {}

func (pb *PanelBase) CancelDrag() {}

func (pb *PanelBase) Tapped() {}

func (pb *PanelBase) Show() {}

func (pb *PanelBase) Hide() {}

func (pb *PanelBase) Visible() bool { return true }

func (pb *PanelBase) Layout(outsideWidth, outsideHeight int) (int, int) {
	if pb.img == nil || outsideWidth != pb.width {
		pb.width = outsideWidth
		pb.img = pb.createImg(BackgroundColor)
		pb.LayoutWidgets()
	}

	return outsideWidth, outsideHeight
}

func (pb *PanelBase) Update() {
	for _, w := range pb.widgets {
		w.Update()
	}
}

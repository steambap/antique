package ui

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

// Label is a button that display some text
type Label struct {
	WidgetBase
	text     string
	fontFace font.Face
	command  string
}

func (l *Label) createImg() *ebiten.Image {
	if l.width == 0 || l.height == 0 {
		return nil
	}

	dc := gg.NewContext(l.width, l.height)
	dc.SetColor(ForegroundColor)
	dc.SetFontFace(l.fontFace)
	dc.DrawString(l.text, 0, float64(l.height)*0.8)

	return ebiten.NewImageFromImage(dc.Image())
}

func measureText(text string, fontFace font.Face) (int, int) {
	dc := gg.NewContext(8, 8)
	dc.SetFontFace(fontFace)
	w, h := dc.MeasureString(text)
	return int(w), int(h)
}

func NewLabel(parent Container, id string, align gg.Align, text string, fontFace font.Face, requestType string) *Label {
	width, height := measureText(text, fontFace)

	l := &Label{
		WidgetBase: WidgetBase{
			parent: parent,
			id:     id,
			img:    nil,
			width:  width,
			height: height,
			align:  align,
		},
		text:     text,
		fontFace: fontFace,
		command:  requestType,
	}
	l.Activate()

	return l
}

func (l *Label) Activate() {
	l.disabled = false
	l.img = l.createImg()
}

func (l *Label) Deactivate() {
	l.disabled = true
	l.img = l.createImg()
}

func (l *Label) Tapped() {
	if l.disabled {
		return
	}
}

func (l *Label) UpdateText(text string) {
	if l.text == text {
		return
	}
	l.text = text
	l.width, l.height = measureText(text, l.fontFace)
	l.img = l.createImg()
	l.parent.LayoutWidgets()
}

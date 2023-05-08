package input

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/steambap/antique/util"
)

// https://gist.github.com/patrickmn/1549985?permalink_comment_id=3036407#gistcomment-3036407
type Observer interface {
	NotifyCallback(StrokeEvent)
}

type Observable interface {
	Add(Observer)
	Remove(Observer)
	Notify(StrokeEvent)
}

// https://ebiten.org/examples/drag.html
type StrokeSource interface {
	Position() (int, int)
	IsJustReleased() bool
}

type MouseStrokeSource struct{}

func (m *MouseStrokeSource) Position() (int, int) {
	return ebiten.CursorPosition()
}

func (m *MouseStrokeSource) IsJustReleased() bool {
	return inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
}

type TouchStrokeSOurce struct {
	ID ebiten.TouchID
}

func (t *TouchStrokeSOurce) Position() (int, int) {
	return ebiten.TouchPosition(t.ID)
}

func (t *TouchStrokeSOurce) IsJustReleased() bool {
	return inpututil.IsTouchJustReleased(t.ID)
}

type WheelStrokeSource struct{}

func (w *WheelStrokeSource) Position() (int, int) {
	x, y := ebiten.Wheel()
	return int(x), int(y)
}

func (w *WheelStrokeSource) IsJustReleased() bool {
	_, y := ebiten.Wheel()
	return y == 0.0
}

type Stroke struct {
	source       StrokeSource
	initX, initY int
	currX, currY int
	released     bool
	observers    sync.Map

	Cancelled     bool
	DraggedObject any
}

func (s *Stroke) Add(observer Observer) {
	s.observers.Store(observer, struct{}{})
}

func (s *Stroke) Remove(observer Observer) {
	s.observers.Delete(observer)
}

func (s *Stroke) Notify(event StrokeEvent) {
	s.observers.Range(func(key, value any) bool {
		if key == nil {
			return false
		}

		key.(Observer).NotifyCallback(event)
		return true
	})
}

func (s *Stroke) IsReleased() bool {
	return s.released
}

func (s *Stroke) Position() (int, int) {
	return s.currX, s.currY
}

func (s *Stroke) PositionDiff() (int, int) {
	return s.currX - s.initX, s.currY - s.initY
}

func NewStroke(source StrokeSource) *Stroke {
	x, y := source.Position()
	return &Stroke{
		source: source,
		initX:  x,
		initY:  y,
		currX:  x,
		currY:  y,
	}
}

type EventType int

const (
	Start EventType = iota + 1
	Move
	Tap
	Stop
	Cancel
)

type StrokeEvent struct {
	Event  EventType
	Stroke *Stroke
	Object any
	X, Y   int
}

func StartStroke(observer Observer) *Stroke {
	var s *Stroke
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s = NewStroke(&MouseStrokeSource{})
	}

	if s == nil {
		ids := []ebiten.TouchID{}
		ids = inpututil.AppendJustPressedTouchIDs(ids)
		if len(ids) > 0 {
			s = NewStroke(&TouchStrokeSOurce{ID: ids[0]})
		}
	}

	if s == nil {
		_, y := ebiten.Wheel()
		if y != 0.0 {
			s = NewStroke(&WheelStrokeSource{})
		}
	}

	if s != nil {
		s.Add(observer)
		s.Notify(StrokeEvent{
			Event:  Start,
			Stroke: s,
			X:      s.initX,
			Y:      s.initY,
		})
	}

	return s
}

func (s *Stroke) Update() {
	if s.released || s.Cancelled {
		return
	}

	if s.source.IsJustReleased() {
		s.released = true
		if util.Abs(s.initX-s.currX) < 4 && util.Abs(s.initY-s.currY) < 4 {
			s.Notify(StrokeEvent{Event: Cancel, Stroke: s, Object: s.DraggedObject, X: s.currX, Y: s.currY})
			s.Notify(StrokeEvent{Event: Tap, Stroke: s, Object: s.DraggedObject, X: s.currX, Y: s.currY})
		} else {
			s.Notify(StrokeEvent{Event: Stop, Stroke: s, Object: s.DraggedObject, X: s.currX, Y: s.currY})
		}
	} else {
		x, y := s.source.Position()
		if s.currX != x || s.currY != y {
			s.currX = x
			s.currY = y
			s.Notify(StrokeEvent{Event: Move, Stroke: s, Object: s.DraggedObject, X: s.currX, Y: s.currY})
		}
	}
}

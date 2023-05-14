package generator

import (
	"github.com/fogleman/gg"
)

// const A = 1
// const TWO_PI = math.Pi * 2

// func NewGalaxy(B, N float64) []gg.Point {
// 	points := make([]gg.Point, 0)
// 	for theta := 0.0; theta < TWO_PI; theta += TWO_PI / 360 {
// 		r := A / math.Log(B*math.Tan(theta/(2*N)))
// 		x := r * math.Cos(theta)
// 		y := r * math.Sin(theta)
// 		points = append(points, gg.Point{X: x, Y: y}, gg.Point{X: -x, Y: -y})
// 	}

// 	return points
// }

func NewGalaxy() []gg.Point {
	return poissonDiskSampling(0.05, 30)
}

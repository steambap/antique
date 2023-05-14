package generator

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/fogleman/gg"
	"github.com/steambap/antique/util"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(123456789))
}

// https://sighack.com/post/poisson-disk-sampling-bridsons-algorithm

func isValidPoint(grid [][]*gg.Point, cellSize float64, nCellWidth int, p gg.Point, radius float64) bool {
	// Make sure the point is on the screen
	if p.X < 0 || p.X >= 1 || p.Y < 0 || p.Y >= 1 {
		return false
	}
	// Check neighbouring cells
	xIndex := (int)(math.Floor(p.X / cellSize))
	yIndex := (int)(math.Floor(p.Y / cellSize))
	i0 := util.Max(xIndex-1, 0)
	i1 := util.Min(xIndex+1, nCellWidth-1)
	j0 := util.Max(yIndex-1, 0)
	j1 := util.Min(yIndex+1, nCellWidth-1)

	for i := i0; i <= i1; i++ {
		for j := j0; j <= j1; j++ {
			existingP := grid[i][j]
			if existingP != nil {
				if p.Distance(*existingP) < radius {
					return false
				}
			}
		}
	}

	return true
}

func poissonDiskSampling(radius float64, k int) []gg.Point {
	points := make([]gg.Point, 0)
	active := make([]gg.Point, 0)

	cellSize := radius / math.Sqrt2
	nCellWidth := int(math.Ceil(1.0/cellSize)) + 1
	fmt.Println(cellSize, nCellWidth)
	var grid [][]*gg.Point = make([][]*gg.Point, nCellWidth)
	for i := 0; i < nCellWidth; i++ {
		grid[i] = make([]*gg.Point, nCellWidth)
	}

	p0 := gg.Point{X: 0.5, Y: 0.5}
	insertPoint(grid, cellSize, p0)
	points = append(points, p0)
	active = append(active, p0)

	for len(active) > 0 {
		randomIndex := r.Intn(len(active))
		p := active[randomIndex]

		found := false

		for tries := 0; tries < k; tries++ {
			theta := float64(r.Intn(360))
			newRadius := (r.Float64() + 1) * radius
			pnewx := p.X + newRadius*math.Cos(gg.Radians(theta))
			pnewy := p.Y + newRadius*math.Sin(gg.Radians(theta))
			pnew := gg.Point{X: pnewx, Y: pnewy}

			if !isValidPoint(grid, cellSize, nCellWidth, pnew, radius) {
				continue
			}

			points = append(points, pnew)
			insertPoint(grid, cellSize, pnew)
			active = append(active, pnew)
			found = true
			break
		}

		if !found {
			active = append(active[:randomIndex], active[randomIndex+1:]...)
		}
	}

	return points
}

func insertPoint(grid [][]*gg.Point, cellSize float64, point gg.Point) {
	xIndex := (int)(math.Floor(point.X / cellSize))
	yIndex := (int)(math.Floor(point.Y / cellSize))
	grid[xIndex][yIndex] = &point
}

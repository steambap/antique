package util

func InRect(x, y int, fn func() (int, int, int, int)) bool {
	x0, y0, x1, y1 := fn()
	return x > x0 && y > y0 && x < x1 && y < y1
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

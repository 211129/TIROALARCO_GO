package models

func Intersect(a, b interface{ Bounds() (float64, float64, float64, float64) }) bool {
	ax1, ay1, ax2, ay2 := a.Bounds()
	bx1, by1, bx2, by2 := b.Bounds()

	return ax1 < bx2 && ax2 > bx1 && ay1 < by2 && ay2 > by1
}
package models

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Arrow struct {
	X, Y   float64
	Speed  float64
	Image  *ebiten.Image
	Hit    bool
	dx, dy float64
}

func NewArrow() *Arrow {
	img, _, err := ebitenutil.NewImageFromFile("./assets/arrow.png")
	if err != nil {
		panic(err)
	}
	return &Arrow{X: 400, Y: 550, Speed: 5, Image: img}
}

func (a *Arrow) Move() {
	if a.Hit {
		return
	}

	a.X += a.dx * a.Speed
	a.Y += a.dy * a.Speed
}

func (a *Arrow) SetDirection(dx, dy float64) {
	a.dx = dx
	a.dy = dy
}

func (a *Arrow) FireStraight() {
	a.dx = 0
	a.dy = -1
}

func (a *Arrow) Bounds() (float64, float64, float64, float64) {
	return a.X, a.Y, a.X + 20, a.Y + 60
}
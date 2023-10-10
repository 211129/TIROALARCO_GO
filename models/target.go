package models

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Target struct {
	X, Y      float64
	Speed     float64
	Direction int
	Image     *ebiten.Image
}

func NewTarget(x float64, y float64) *Target {
	img, _, err := ebitenutil.NewImageFromFile("assets/target.png")
	if err != nil {
		panic(err)
	}
	return &Target{X: x, Y: y, Speed: 2, Direction: 1, Image: img}
}

func (t *Target) Move() {
	t.X += t.Speed * float64(t.Direction)
	if t.X > 780 || t.X < 20 {
		t.Direction *= -1
	}
}

func (t *Target) Bounds() (float64, float64, float64, float64) {
	return t.X, t.Y, t.X + 100, t.Y + 100
}
package models

import (
	"math"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Arrow struct {
	X, Y   float64
	Speed  float64
	Image  *ebiten.Image
	Hit    bool
}

func NewArrow() *Arrow {
	img, _, err := ebitenutil.NewImageFromFile("./assets/arrow.png")
	if err != nil {
		panic(err)
	}
	return &Arrow{X: 400, Y: 550, Speed: 5, Image: img}
}

func (a *Arrow) Move(targets []*Target) {
	if a.Hit {  // Si la flecha ya ha golpeado un target, no la movemos más.
		return
	}
	closestTarget := targets[0]
	closestDistance := math.Sqrt(math.Pow(closestTarget.X-a.X, 2) + math.Pow(closestTarget.Y-a.Y, 2))

	for _, target := range targets {
		distance := math.Sqrt(math.Pow(target.X-a.X, 2) + math.Pow(target.Y-a.Y, 2))
		if distance < closestDistance {
			closestTarget = target
			closestDistance = distance
		}
	}

	if closestTarget.X > a.X {
		a.X += a.Speed
	} else if closestTarget.X < a.X {
		a.X -= a.Speed
	}

	if closestTarget.Y > a.Y {
		a.Y += a.Speed
	} else if closestTarget.Y < a.Y {
		a.Y -= a.Speed
	}
}

func (a *Arrow) Bounds() (float64, float64, float64, float64) {
	return a.X, a.Y, a.X + 20, a.Y + 60
}
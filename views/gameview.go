package views

import (
	"fmt"
	"TIROALARCO/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameView struct {
	Arrows             []*models.Arrow
	Targets            []*models.Target
	Score              int
	SpacebarLastFrame  bool
}

func NewGameView() *GameView {
	return &GameView{
		Arrows:  []*models.Arrow{},
		Targets: []*models.Target{models.NewTarget(100, 100), models.NewTarget(600, 150)},
	}
}

func (v *GameView) Update() {
	
	spacebarPressed := ebiten.IsKeyPressed(ebiten.KeySpace)
	if spacebarPressed && !v.SpacebarLastFrame {
		newArrow := models.NewArrow()
		newArrow.FireStraight()
		v.Arrows = append(v.Arrows, newArrow)
	}
	v.SpacebarLastFrame = spacebarPressed

	
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		for _, arrow := range v.Arrows {
			arrow.SetDirection(-1, 0)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		for _, arrow := range v.Arrows {
			arrow.SetDirection(1, 0)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		for _, arrow := range v.Arrows {
			arrow.SetDirection(0, -1)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		for _, arrow := range v.Arrows {
			arrow.SetDirection(0, 1)
		}
	}

	for _, arrow := range v.Arrows {
		arrow.Move()
	}

	for _, arrow := range v.Arrows {
		if arrow.Hit {
			continue
		}
		for _, target := range v.Targets {
			if models.Intersect(arrow, target) {
				arrow.Hit = true
				v.Score++
				break
			}
		}
	}

	v.Arrows = removeHitArrows(v.Arrows)
	for _, target := range v.Targets {
		target.Move()
	}
}

func removeHitArrows(arrows []*models.Arrow) []*models.Arrow {
	newArrows := []*models.Arrow{}
	for _, arrow := range arrows {
		if !arrow.Hit {
			newArrows = append(newArrows, arrow)
		}
	}
	return newArrows
}

func (v *GameView) Draw(screen *ebiten.Image) {
	bg, _, err := ebitenutil.NewImageFromFile("assets/paris.png")
	if err != nil {
		panic(err)
	}
	screen.DrawImage(bg, &ebiten.DrawImageOptions{})

	for _, arrow := range v.Arrows {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(arrow.X, arrow.Y)
		screen.DrawImage(arrow.Image, opts)
	}
	for _, target := range v.Targets {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(target.X, target.Y)
		screen.DrawImage(target.Image, opts)
	}

	scoreText := fmt.Sprintf("Puntaje: %d", v.Score)
	ebitenutil.DebugPrint(screen, scoreText)
}

func (v *GameView) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

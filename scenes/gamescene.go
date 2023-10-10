package scenes

import (
	"TIROALARCO/views"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	view *views.GameView
}

func NewGameScene() *GameScene {
	return &GameScene{
		view: views.NewGameView(),
	}
}

func (g *GameScene) Update() error {
	g.view.Update()
	return nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	g.view.Draw(screen)
}

func (g *GameScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}
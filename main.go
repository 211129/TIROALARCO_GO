package main

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"TIROALARCO/scenes"
)

func main() {
	gameScene := scenes.NewGameScene()
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Tiro al Arco Paris 2024")
	if err := ebiten.RunGame(gameScene); err != nil {
		log.Fatal(err)
	}
}

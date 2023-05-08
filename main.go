package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/steambap/antique/antique"
)

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)
	ebiten.SetWindowTitle("Antique")

	antique.NewGame()
	if err := ebiten.RunGame(antique.TheGame); err != nil {
		log.Fatal(err)
	}
}

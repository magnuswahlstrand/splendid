package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/splendid/game"
)

func main() {

	g := game.New()
	if err := ebiten.Run(g.Update, game.Size.Width, game.Size.Height, 1, "my application"); err != nil {
		log.Fatal(err)
	}
}

package game

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/splendid/internal/input"
)

type gameState string

const (
	stopped  gameState = ""
	running  gameState = "running"
	paused   gameState = "paused"
	finished gameState = "finished"
)

// Update is the main game loop
func (g *Game) Update(screen *ebiten.Image) error {

	switch g.state {
	case running:
		// Change state
		if g.TimeRemaining() < 0 {
			g.state = finished
			return nil
		}

		g.draw(screen)
		if input.MouseJustPressed() {
			g.handleClick()
		}

	case paused:
		g.drawPaused(screen)
	case finished:
		g.drawFinished(screen)
		if input.MouseJustPressed() && time.Since(g.startTime) > MinTimeBeforeNextGame {
			g.Restart()
		}

	case stopped:
		panic("Mo")
	}

	return nil
}

func (g *Game) handleClick() {
	cursorPos := input.MousePosition()
	cardPressed, correctPressed := g.board.CardPressedCorrectPressed(cursorPos)

	if cardPressed {
		if correctPressed {
			fmt.Println("correct card pressed")
			g.score++
		} else {
			fmt.Println("incorrect card pressed")
			g.vibrate()
		}
		g.board.NewRound(CardsOnBoard)
	}
}

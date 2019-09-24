package game

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/splendid/internal/input"
)

type gameState string

const (
	stopped   gameState = ""
	highscore gameState = "highscore"
	running   gameState = "running"
	paused    gameState = "paused"
	finished  gameState = "finished"
)

// Update is the main game loop
func (g *Game) Update(screen *ebiten.Image) error {

	switch g.state {
	case highscore:
		g.drawHighscore(screen)
		if input.MouseJustPressed() && time.Since(g.startTime) > MinTimeBeforeNextGame {
			g.Restart()
		}

	case running:
		// Change state
		if g.TimeRemaining() < 0 {
			g.state = highscore
			g.startTime = time.Now()
			go g.postHighscore()
			return nil
		}

		g.draw(screen)
		if input.MouseJustPressed() {
			g.handleClick()
		}

	case paused:
		g.drawPaused(screen)
	// case finished:
	// 	g.drawFinished(screen)
	// 	if input.MouseJustPressed() && time.Since(g.startTime) > MinTimeBeforeNextGame {
	// 		g.Restart()
	// 	}

	default:
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
			g.mobileConnector.Vibrate()
		}
		g.board.NewRound(CardsOnBoard)
	}
}

func (g *Game) postHighscore() {
	// Todo, fix this in a better way
	if g.client.GetClientID() == "" {
		g.client.SetClientID(g.mobileConnector.GetDeviceID())
	}

	// Save score to avoid race conditions

	score := g.score
	go func() {
		if err := g.client.AddSimple(float64(score)); err != nil {
			println("failed to save score to server" + err.Error())
		}
		g.updateHighscore()
	}()
}

func (g *Game) updateHighscore() {
	list, err := g.client.ListSimple()
	if err != nil {
		println("failed to load highscore" + err.Error())
		return
	}

	// Todo, do this safely
	g.highscore = list
}

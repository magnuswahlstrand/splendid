package game

import (
	"fmt"

	"github.com/peterhellberg/gfx"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/splendid/card"
	"github.com/kyeett/splendid/font"
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
		if input.MouseJustPressed() {
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
		g.board.NewRound(4)
	}
}

var (
	pauseScreen, pauseScreen2 *ebiten.Image
)

func init() {
	var err error

	pauseScreen, err = ebiten.NewImage(Size.Width, Size.Height, ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	pauseScreen.Fill(colornames.Gray)
}

func (g *Game) draw(screen *ebiten.Image) {
	// Draw score and time
	boardHeight := g.board.Height()
	scoreBox := gfx.BoundsToRect(gfx.IR(0, boardHeight, card.DrawWidth, boardHeight+100))
	font.DrawInCenter(screen, fmt.Sprintf("Score: %d", g.score), scoreBox, colornames.White)
	timeBox := gfx.BoundsToRect(gfx.IR(0, boardHeight, card.DrawWidth, boardHeight+100)).Moved(gfx.IV(card.DrawWidth+card.DrawPadding, 0))
	font.DrawInCenter(screen, fmt.Sprintf("Time: %0.1f", g.TimeRemaining().Seconds()), timeBox, colornames.White)

	// Draw board
	g.board.Draw(screen)
}

func (g *Game) drawPaused(screen *ebiten.Image) {
	g.draw(screen)

	// Draw a semi-transparent screen
	opt := &ebiten.DrawImageOptions{}
	opt.ColorM.Scale(1, 1, 1, 0.9)
	screen.DrawImage(pauseScreen, opt)
}

func (g *Game) drawFinished(screen *ebiten.Image) {
	line1Box := gfx.R(0, 0, float64(Size.Width), 100).Moved(gfx.IV(0, Size.Height/2-100))
	font.DrawInCenterBig(screen, "Game", line1Box, colornames.White)
	line2Box := gfx.R(0, 0, float64(Size.Width), 100).Moved(gfx.IV(0, Size.Height/2-40))
	font.DrawInCenterBig(screen, "finished", line2Box, colornames.White)

	scoreBox := gfx.R(0, 0, float64(Size.Width), 30).Moved(gfx.IV(0, Size.Height/2+55))
	font.DrawInCenter(screen, fmt.Sprintf("(score: %d)", g.score), scoreBox, colornames.White)

}

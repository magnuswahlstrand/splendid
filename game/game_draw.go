package game

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/splendid/card"
	"github.com/kyeett/splendid/font"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
)

var (
	pauseScreen *ebiten.Image
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
	line1Box := gfx.R(0, 0, float64(Size.Width), 100).Moved(gfx.IV(0, Size.Height/2-130))
	font.DrawInCenterBig(screen, "Game", line1Box, colornames.White)
	line2Box := gfx.R(0, 0, float64(Size.Width), 100).Moved(gfx.IV(0, Size.Height/2-70))
	font.DrawInCenterBig(screen, "finished", line2Box, colornames.White)

	scoreBox := gfx.R(0, 0, float64(Size.Width), 30).Moved(gfx.IV(0, Size.Height/2+25))
	font.DrawInCenter(screen, fmt.Sprintf("(score: %d)", g.score), scoreBox, colornames.White)

	// Wait a few seconds before showing continue text
	if time.Since(g.startTime) > MinTimeBeforeNextGame {

		// Blink every 500 ms
		showText := (2*time.Since(g.startTime)/(time.Second))%2 == 0
		if showText {
			continueBox := gfx.R(0, 0, float64(Size.Width), 30).Moved(gfx.IV(0, Size.Height/2+100))
			font.DrawInCenter(screen, "click to continue", continueBox, colornames.Lightgray)
		}
	}

}

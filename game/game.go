package game

import (
	"time"

	"github.com/kyeett/splendid/card"
	"github.com/peterhellberg/gfx"
)

var Size = struct {
	Width  int
	Height int
}{
	Width:  128*2 + 20,
	Height: 128*2 + 20 + 200,
}

type Game struct {
	state       gameState
	score       int
	cardHint    string
	endTime     time.Time
	correctCard *card.Card
	board       card.Board
	Vibrator    func()
}

func New() *Game {
	g := &Game{
		board: card.Board{
			Position: gfx.V(0, 0),
		},
	}

	g.Restart()
	return g
}

func (g *Game) Restart() {
	g.state = running
	g.score = 0
	g.board.NewRound(4)
	g.endTime = time.Now().Add(1 * time.Second)
}

func (g *Game) TimeRemaining() time.Duration {
	return g.endTime.Sub(time.Now())
}

func (g *Game) Layout(viewWidth, viewHeight int) (screenWidth, screenHeight int) {
	return Size.Width, Size.Height
}
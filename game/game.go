package game

import (
	"time"

	"github.com/kyeett/highscore-server/client"
	highscoreModel "github.com/kyeett/highscore-server/model"

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
	state           gameState
	score           int
	client          client.HighscoreClient
	highscore       []*highscoreModel.Score
	cardHint        string
	startTime       time.Time
	correctCard     *card.Card
	board           card.Board
	mobileConnector MobileConnector
}

func New(lateInit bool) *Game {

	var c client.HighscoreClient
	var err error
	if lateInit {
		c, err = client.NewNoID(HighscoreServerURL, GameSplendidName)
		if err != nil {
			println("failed to create highscore client" + err.Error())
		}
	} else {
		c, err = client.New(HighscoreServerURL, GameSplendidName)
		if err != nil {
			println("failed to create highscore client" + err.Error())
		}
	}

	g := &Game{
		client: c,
		board: card.Board{
			Position: gfx.V(0, 0),
		},
		mobileConnector: &dummyMobileConnector{},
	}

	g.Restart()
	go g.updateHighscore()
	g.state = highscore
	return g
}

func (g *Game) Restart() {
	g.state = running
	g.score = 0
	g.board.NewRound(CardsOnBoard)
	g.startTime = time.Now()
	g.highscore = nil // Prepare for next time
}

func (g *Game) TimeRemaining() time.Duration {
	return g.startTime.Add(GameDuration * time.Second).Sub(time.Now())
}

func (g *Game) Layout(viewWidth, viewHeight int) (screenWidth, screenHeight int) {
	return Size.Width, Size.Height
}

func (g *Game) SetMobileConnector(mc MobileConnector) {
	g.mobileConnector = mc
}

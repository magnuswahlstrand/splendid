package splendidmobile

import (
	"github.com/hajimehoshi/ebiten/mobile"
	"github.com/kyeett/splendid/game"
)

// Layout is called when the game is initialized or the view size is changed.
// You can return a fixed screen size, or you can calculate a screen size based on the given view size.
// The scaling is automatically adjusted.

var g *game.Game

func init() {
	g = game.New(true)
	mobile.SetGame(g)
}

func SetMobileConnector(mc GameMobileInterface) {
	g.SetMobileConnector(mc)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}

// type Counter struct {
// 	Value int
// }

// func (c *Counter) Inc() { c.Value++ }

// func NewCounter() *Counter { return &Counter{5} }

// type Printer interface {
// 	Print(s string)
// }

// func PrintHello(p Printer) {
// 	p.Print("Hello, World!")
// }

type GameMobileInterface interface {
	Vibrate()
	ShowEndDialog()
	GetDeviceID() string
}

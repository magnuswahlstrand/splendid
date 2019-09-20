package card

import (
	"math/rand"
	"strings"

	"github.com/kyeett/splendid/font"

	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/gfx"
)

const (
	columns     = 2
	cardOffsetY = 100
)

type Board struct {
	Position   gfx.Vec
	Dimensions gfx.Rect

	hint        string
	correctCard *Card
	Cards       []*Card
}

func (b *Board) Width() int {
	return columns*DrawWidth + (columns-1)*DrawPadding
}

func (b *Board) Height() int {
	rows := (len(b.Cards) + 1) / columns
	return rows*DrawHeight + (rows-1)*DrawPadding + cardOffsetY
}

func (b *Board) NewRound(n int) {
	b.Cards = RandomCardSlice(n)
	b.correctCard = b.Cards[rand.Intn(n)]
	b.hint = b.correctCard.ColorOrTypeHint()
}

func (b *Board) Draw(screen *ebiten.Image) {
	for i, card := range b.Cards {
		col := i % columns
		row := i / columns
		pos := b.Position.AddXY(float64(col)*(DrawPadding+DrawWidth), float64(row)*(DrawPadding+DrawWidth)+cardOffsetY)
		card.draw(screen, pos)
	}

	// Draw hint text
	box := gfx.R(0, 0, DrawPadding+2*DrawWidth, 100).Moved(gfx.V(b.Position.X, 0))
	font.DrawInCenterBig(screen, strings.ToUpper(b.hint), box, colornames.White)
}

func (b *Board) CardPressedCorrectPressed(cursorPos gfx.Vec) (bool, bool) {
	pressedCard := b.PressedCard(cursorPos)
	return pressedCard != nil, pressedCard == b.correctCard
}

func (b *Board) PressedCard(cursorPos gfx.Vec) *Card {
	for i, card := range b.Cards {
		col := i % columns
		row := i / columns
		pos := b.Position.AddXY(float64(col)*(DrawPadding+DrawWidth), float64(row)*(DrawPadding+DrawWidth)+cardOffsetY)
		bounds := gfx.R(0, 0, DrawWidth, DrawHeight).Moved(pos)
		if bounds.Contains(cursorPos) {
			return card
		}
	}
	return nil
}

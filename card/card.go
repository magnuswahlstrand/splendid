package card

import (
	"math/rand"
	"time"
)

// Type of cards
type Type string

const (
	kraken Type = "kraken"
	pig    Type = "pig"
	rose   Type = "rose"
	stag   Type = "stag"
	wolf   Type = "wolf"
)

var types = []Type{kraken, pig, rose, stag, wolf}

// Card object consists of a Type and a color
type Card struct {
	Type  Type
	Color cardColor
}

// ColorOrTypeHint returns card color or card type as a string
func (c Card) ColorOrTypeHint() string {
	if rand.Intn(2) == 0 {
		return c.Color.String()
	}
	return string(c.Type)
}

// RandomCardSlice returns a slice of n random cards.
// Each type and color only used once
func RandomCardSlice(n int) []*Card {
	// Shuffle cards types and colors
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(colors), func(i, j int) { colors[i], colors[j] = colors[j], colors[i] })
	rand.Shuffle(len(types), func(i, j int) { types[i], types[j] = types[j], types[i] })

	var cards []*Card
	for i := 0; i < n; i++ {
		cards = append(cards, &Card{
			Type:  types[i],
			Color: colors[i],
		})
	}
	return cards
}

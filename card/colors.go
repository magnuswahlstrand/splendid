package card

import (
	"image/color"
)

type cardColor color.RGBA

var (
	colorRed    = cardColor(color.RGBA{0xbb, 0x57, 0x35, 0xFF})
	colorOrange = cardColor(color.RGBA{0xdf, 0x92, 0x45, 0xFF})
	colorYellow = cardColor(color.RGBA{0xec, 0xd2, 0x74, 0xFF})
	colorGreen  = cardColor(color.RGBA{0x83, 0xa8, 0x16, 0xFF})
	colorBlue   = cardColor(color.RGBA{0x04, 0x68, 0x94, 0xFF})
	colorWhite  = cardColor(color.RGBA{0xfd, 0xf9, 0xf1, 0xFF})
	colorBeige  = cardColor(color.RGBA{0xc7, 0xb2, 0x95, 0xFF})
)

// var (
// 	colorRed    = cardColor(color.RGBA{208, 2, 27, 255})
// 	colorPurple = cardColor(color.RGBA{189, 16, 224, 255})
// 	colorGreen  = cardColor(color.RGBA{126, 211, 33, 255})
// 	colorBlue   = cardColor(color.RGBA{72, 186, 255, 255})
// 	colorWhite  = cardColor(color.RGBA{255, 255, 255, 255})
// 	colorYellow = cardColor(color.RGBA{248, 231, 28, 255})
// )

var colors = []cardColor{
	colorRed,
	colorOrange,
	colorYellow,
	colorGreen,
	colorBlue,
	colorWhite,
	colorBeige,
}

func (c cardColor) String() string {
	switch c {
	case colorRed:
		return "red"
	case colorOrange:
		return "orange"
	case colorYellow:
		return "yellow"
	case colorGreen:
		return "green"
	case colorBlue:
		return "blue"
	case colorWhite:
		return "white"
	case colorBeige:
		return "beige"
	default:
		return "invalid"
	}
}

func (c cardColor) RScale() float64 {
	return float64(c.R) / 256
}
func (c cardColor) GScale() float64 {
	return float64(c.G) / 256
}
func (c cardColor) BScale() float64 {
	return float64(c.B) / 256
}

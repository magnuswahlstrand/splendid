package ebitendrawutil

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/ebitenutil"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/gfx"
)

// DrawRect draws a rectangle of color clr at on the bounds of gfx.Rect r
// If no width is specified, 1 will be used
// If one or more values are specified, the first value will be used as width
func DrawRect(screen *ebiten.Image, r gfx.Rect, clr color.Color, width ...int) {
	thickness := 1.0
	if len(width) > 0 {
		thickness = float64(width[0])
	}

	max := r.Max.Sub(gfx.V(thickness-1, thickness-1))
	ebitenutil.DrawRect(screen, r.Min.X, r.Min.Y, r.W(), thickness, clr)
	ebitenutil.DrawRect(screen, r.Min.X, r.Min.Y, thickness, r.H(), clr)
	ebitenutil.DrawRect(screen, r.Min.X, max.Y, r.W(), thickness, clr)
	ebitenutil.DrawRect(screen, max.X, r.Min.Y, thickness, r.H(), clr)
}

// ImageFromBytes takes a byte array of an PNG image and returns an *ebiten.Image
// If the bytes are not valid PNG, it will log.Fatal(err)
func ImageFromBytes(b []byte) *ebiten.Image {
	img, err := gfx.DecodeImage(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}

	eImg, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	return eImg
}

func BoundingBoxFromString(s string, fnt font.Face) gfx.Rect {
	width := font.MeasureString(fnt, s).Ceil()
	height := fnt.Metrics().Height.Ceil()
	return gfx.R(0, 0, float64(width), float64(height))
}

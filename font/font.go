package font

import (
	"fmt"
	"image/color"
	"log"

	"github.com/peterhellberg/gfx"

	"github.com/kyeett/ebitendrawutil"

	"github.com/hajimehoshi/ebiten"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

var (
	MPlusNormalFont font.Face
	MPlusBigFont    font.Face
)

func init() {
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	MPlusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    20,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	MPlusBigFont = truetype.NewFace(tt, &truetype.Options{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func Draw(screen *ebiten.Image, s string, x, y int, clr color.Color) {
	text.Draw(screen, s, MPlusNormalFont, x, y, clr)
}

func DrawBig(screen *ebiten.Image, s string, x, y int, clr color.Color) {
	text.Draw(screen, s, MPlusBigFont, x, y, clr)
}

func DrawInCenter(screen *ebiten.Image, s string, r gfx.Rect, clr color.Color) {
	drawInCenter(screen, s, r, MPlusNormalFont, clr)
}

func DrawInCenterBig(screen *ebiten.Image, s string, r gfx.Rect, clr color.Color) {
	drawInCenter(screen, s, r, MPlusBigFont, clr)
}

func drawInCenter(screen *ebiten.Image, s string, r gfx.Rect, fnt font.Face, clr color.Color) {
	bb := ebitendrawutil.BoundingBoxFromString(s, fnt)

	offsetX := (r.W() - bb.W()) / 2
	offsetY := (r.H() - bb.H()) / 2

	x := int(r.Min.X + offsetX)
	y := int(r.Min.Y + offsetY + bb.H())
	fmt.Println(offsetX, offsetY)
	fmt.Println(x, y)

	text.Draw(screen, s, fnt, x, y, clr)
	// ebitendrawutil.DrawRect(screen, r, colornames.Green, 2)
}

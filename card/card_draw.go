package card

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kyeett/splendid/assets"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
)

const (
	DrawHeight  = 128
	DrawWidth   = 128
	DrawPadding = 20
)

var images map[Type]*ebiten.Image

func init() {
	images = map[Type]*ebiten.Image{
		kraken: loadEbitenImg("img/kraken.png"),
		pig:    loadEbitenImg("img/pig.png"),
		rose:   loadEbitenImg("img/rose.png"),
		stag:   loadEbitenImg("img/stag.png"),
		wolf:   loadEbitenImg("img/wolf.png"),
	}
}

func (c Card) draw(screen *ebiten.Image, pos gfx.Vec) {
	// Background
	ebitenutil.DrawRect(screen, pos.X, pos.Y, DrawWidth, DrawHeight, colornames.Gray)

	opt := ebiten.DrawImageOptions{}
	opt.GeoM.Translate(pos.X, pos.Y)
	opt.ColorM.Scale(c.Color.RScale(), c.Color.GScale(), c.Color.BScale(), 1)
	screen.DrawImage(images[c.Type], &opt)
}

func loadEbitenImg(path string) *ebiten.Image {
	// img, err := gfx.OpenImage(path)
	img, err := gfx.DecodeImageBytes(assets.MustAsset(path))
	if err != nil {
		panic(err)
	}
	ebitenImg, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		panic(err)
	}
	return ebitenImg
}

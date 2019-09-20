package input

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/peterhellberg/gfx"
)

func MousePosition() gfx.Vec {
	if len(ebiten.TouchIDs()) > 0 {
		ID := ebiten.TouchIDs()[0]
		x, y := ebiten.TouchPosition(ID)
		return gfx.V(float64(x), float64(y))
	}

	x, y := ebiten.CursorPosition()
	return gfx.V(float64(x), float64(y))
}

func MouseJustPressed() bool {
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || len(inpututil.JustPressedTouchIDs()) > 0
}

func MousePressed() bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) || len(ebiten.TouchIDs()) > 0
}

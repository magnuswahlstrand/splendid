package game

func (g *Game) vibrate() {
	if g.Vibrator == nil {
		return
	}

	g.Vibrator()
}

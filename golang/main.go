package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(scr *ebiten.Image) {
	scr.Set(160, 120, color.White)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return 320, 160
}

func main() {
	g := &Game{}
	ebiten.RunGame(g)
}

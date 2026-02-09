package virus

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	gameOverImg *ebiten.Image
)

func (g *Game) DrawGameOverScreen(screen *ebiten.Image) {
	screen.DrawImage(gameOverImg, nil)
	text := "GAME OVER"
	textSize := 60
	textWidth := len(text) * textSize / 5
	x := (screen.Bounds().Dx() - textWidth) / 2
	y := screen.Bounds().Dy() / 2
	ebitenutil.DebugPrintAt(screen, text, x, y)

	subText := "PRESS R TO RESTART"
	subTextSize := 30
	subTextWidth := len(subText) * subTextSize / 5
	subX := (screen.Bounds().Dx() - subTextWidth) / 2
	subY := y + textSize
	ebitenutil.DebugPrintAt(screen, subText, subX, subY)

	gradeText := "GRADE:" + Grade
	gradeTextSize := 10
	gradeTextWidth := len(gradeText) * gradeTextSize / 5
	gradeX := (screen.Bounds().Dx() - gradeTextWidth) / 2
	gradeY := subY + subTextSize + 20
	ebitenutil.DebugPrintAt(screen, gradeText, gradeX, gradeY)
}

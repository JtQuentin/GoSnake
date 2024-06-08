package game

import (
	"fmt"
	"image/color"
	"log"

	"GoSnake/food"
	"GoSnake/vars"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

type Renderer struct {
	screen *ebiten.Image
	face   font.Face
}

func NewRenderer() *Renderer {
	return &Renderer{
		face: basicfont.Face7x13,
	}
}

func (r *Renderer) drawBackground() {
	r.screen.Fill(color.RGBA{154, 198, 0, 255})
}

func (r *Renderer) drawSnake(body []food.Point) {
	for _, p := range body {
		ebitenutil.DrawRect(r.screen, float64(p.X*vars.TileSize), float64(p.Y*vars.TileSize), vars.TileSize, vars.TileSize, color.RGBA{33, 50, 15, 255})
	}
}

func (r *Renderer) drawFood(position food.Point) {
	ebitenutil.DrawRect(r.screen, float64(position.X*vars.TileSize), float64(position.Y*vars.TileSize), vars.TileSize, vars.TileSize, color.RGBA{231, 71, 29, 255})
}

func (r *Renderer) drawUI(score int, gameOver bool, gameWon bool, gameStarted bool, gamePaused bool) {
	scoreText := fmt.Sprintf("Score: %d", score)
	text.Draw(r.screen, scoreText, r.face, 5, vars.ScreenHeight-5, color.White)

	if !gameStarted {
		startText := "Press 'SPACE' to start the game"
		startTextWidth := text.BoundString(r.face, startText).Dx()
		x := (vars.ScreenWidth - startTextWidth) / 2
		text.Draw(r.screen, startText, r.face, x, vars.ScreenHeight/2, color.White)
	} else {
		if gameOver {
			gameOverText := "Game Over"
			gameOverTextWidth := text.BoundString(r.face, gameOverText).Dx()
			x := (vars.ScreenWidth - gameOverTextWidth) / 2
			text.Draw(r.screen, gameOverText, r.face, x, vars.ScreenHeight/2, color.White)

			restartText := "Press 'R' to restart"
			restartTextWidth := text.BoundString(r.face, restartText).Dx()
			x = (vars.ScreenWidth - restartTextWidth) / 2
			text.Draw(r.screen, restartText, r.face, x, vars.ScreenHeight/2+16, color.White)
			scores, err := LoadScores()
			if err == nil {
				startY := vars.ScreenHeight/2 + 32
				for i, entry := range scores {
					if i >= 5 {
						break
					}
					scoreLine := fmt.Sprintf("%d. %s: %d", i+1, entry.Name, entry.Score)
					text.Draw(r.screen, scoreLine, r.face, vars.ScreenWidth/2-60, startY+(i*16), color.White)
				}
			} else {
				log.Printf("Error loading scores: %v", err)
			}
		}

		if gameWon {
			gameOverText := "You Won!"
			gameOverTextWidth := text.BoundString(r.face, gameOverText).Dx()
			x := (vars.ScreenWidth - gameOverTextWidth) / 2
			text.Draw(r.screen, gameOverText, r.face, x, vars.ScreenHeight/2, color.White)

			restartText := "Press 'R' to restart"
			restartTextWidth := text.BoundString(r.face, restartText).Dx()
			x = (vars.ScreenWidth - restartTextWidth) / 2
			text.Draw(r.screen, restartText, r.face, x, vars.ScreenHeight/2+16, color.White)
		}

		if gamePaused {
			pausedText := "You paused the game"
			pausedTextWidth := text.BoundString(r.face, pausedText).Dx()
			x := (vars.ScreenWidth - pausedTextWidth) / 2
			text.Draw(r.screen, pausedText, r.face, x, vars.ScreenHeight/2-16, color.White)

			resumeText := "Press 'P' to resume"
			resumeTextWidth := text.BoundString(r.face, resumeText).Dx()
			x = (vars.ScreenWidth - resumeTextWidth) / 2
			text.Draw(r.screen, resumeText, r.face, x, vars.ScreenHeight/2, color.White)
		}
	}
}

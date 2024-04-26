package main

import (
	"fmt"
	"image/color"
	"log"

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

func (r *Renderer) drawSnake(body []Point) {
	for _, p := range body {
		ebitenutil.DrawRect(r.screen, float64(p.X*tileSize), float64(p.Y*tileSize), tileSize, tileSize, color.RGBA{33, 50, 15, 255})
	}
}

func (r *Renderer) drawFood(position Point) {
	ebitenutil.DrawRect(r.screen, float64(position.X*tileSize), float64(position.Y*tileSize), tileSize, tileSize, color.RGBA{231, 71, 29, 255})
}

func (r *Renderer) drawUI(score int, gameOver bool, gameWon bool, gameStarted bool) {
	scoreText := fmt.Sprintf("Score: %d", score)
	text.Draw(r.screen, scoreText, r.face, 5, screenHeight-5, color.White)

	if !gameStarted {
		text.Draw(r.screen, "Press 'SPACE' to start the game", r.face, screenWidth/2-100, screenHeight/2, color.White)
	} else {
		if gameOver {
			text.Draw(r.screen, "Game Over", r.face, screenWidth/2-40, screenHeight/2, color.White)
			text.Draw(r.screen, "Press 'R' to restart", r.face, screenWidth/2-60, screenHeight/2+16, color.White)
			scores, err := LoadScores()
			if err == nil {
				startY := screenHeight/2 + 32
				for i, entry := range scores {
					if i >= 5 {
						break
					}
					scoreLine := fmt.Sprintf("%d. %s: %d", i+1, entry.Name, entry.Score)
					text.Draw(r.screen, scoreLine, r.face, screenWidth/2-60, startY+(i*16), color.White)
				}
			} else {
				log.Printf("Error loading scores: %v", err)
			}
		}

		if gameWon {
			text.Draw(r.screen, "You Won!", r.face, screenWidth/2-40, screenHeight/2, color.White)
			text.Draw(r.screen, "Press 'R' to restart", r.face, screenWidth/2-60, screenHeight/2+16, color.White)
		}
	}
}

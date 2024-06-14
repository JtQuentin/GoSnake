package game

import (
	"fmt"
	"image/color"
	"log"

	"GoSnake/vars"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

// Renderer handles rendering the game
type Renderer struct {
	screen *ebiten.Image // The screen image to render on
	face   font.Face     // The font face to use for rendering text
}

// NewRenderer creates a new Renderer instance
func NewRenderer() *Renderer {
	return &Renderer{
		face: basicfont.Face7x13, // Using a basic font face
	}
}

// drawBackground fills the screen with a green color
func (r *Renderer) drawBackground() {
	r.screen.Fill(color.RGBA{154, 198, 0, 255})
}

// drawSnake draws the snake's body on the screen
func (r *Renderer) drawSnake(body []vars.Point) {
	for _, p := range body {
		ebitenutil.DrawRect(r.screen, float64(p.X*vars.TileSize), float64(p.Y*vars.TileSize), vars.TileSize, vars.TileSize, color.RGBA{33, 50, 15, 255})
	}
}

// drawFood draws the food on the screen
func (r *Renderer) drawFood(position vars.Point) {
	ebitenutil.DrawRect(r.screen, float64(position.X*vars.TileSize), float64(position.Y*vars.TileSize), vars.TileSize, vars.TileSize, color.RGBA{231, 71, 29, 255})
}

// drawUI draws the user interface elements on the screen
func (r *Renderer) drawUI(score int, gameOver bool, gameWon bool, gameStarted bool, gamePaused bool) {
	// Draw the score
	scoreText := fmt.Sprintf("Score: %d", score)
	text.Draw(r.screen, scoreText, r.face, 5, vars.ScreenHeight-5, color.White)

	// Draw the start game text if the game has not started
	if !gameStarted {
		startText := "Press 'SPACE' to start the game"
		startTextWidth := text.BoundString(r.face, startText).Dx()
		x := (vars.ScreenWidth - startTextWidth) / 2
		text.Draw(r.screen, startText, r.face, x, vars.ScreenHeight/2, color.White)
	} else {
		// Draw game over text and restart instructions if the game is over
		if gameOver {
			// Draw game over text
			gameOverText := "Game Over"
			gameOverTextWidth := text.BoundString(r.face, gameOverText).Dx()
			x := (vars.ScreenWidth - gameOverTextWidth) / 2
			text.Draw(r.screen, gameOverText, r.face, x, vars.ScreenHeight/2, color.White)

			// Draw restart instructions
			restartText := "Press 'R' to restart"
			restartTextWidth := text.BoundString(r.face, restartText).Dx()
			x = (vars.ScreenWidth - restartTextWidth) / 2
			text.Draw(r.screen, restartText, r.face, x, vars.ScreenHeight/2+16, color.White)

			// Draw the high scores
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

		// Draw game won text and restart instructions if the game is won
		if gameWon {
			// Draw game won text
			gameOverText := "You Won!"
			gameOverTextWidth := text.BoundString(r.face, gameOverText).Dx()
			x := (vars.ScreenWidth - gameOverTextWidth) / 2
			text.Draw(r.screen, gameOverText, r.face, x, vars.ScreenHeight/2, color.White)

			// Draw restart instructions
			restartText := "Press 'R' to restart"
			restartTextWidth := text.BoundString(r.face, restartText).Dx()
			x = (vars.ScreenWidth - restartTextWidth) / 2
			text.Draw(r.screen, restartText, r.face, x, vars.ScreenHeight/2+16, color.White)
		}

		// Draw paused game text and resume instructions if the game is paused
		if gamePaused {
			// Draw paused game text
			pausedText := "You paused the game"
			pausedTextWidth := text.BoundString(r.face, pausedText).Dx()
			x := (vars.ScreenWidth - pausedTextWidth) / 2
			text.Draw(r.screen, pausedText, r.face, x, vars.ScreenHeight/2-16, color.White)

			// Draw resume instructions
			resumeText := "Press 'P' to resume"
			resumeTextWidth := text.BoundString(r.face, resumeText).Dx()
			x = (vars.ScreenWidth - resumeTextWidth) / 2
			text.Draw(r.screen, resumeText, r.face, x, vars.ScreenHeight/2, color.White)
		}
	}
}

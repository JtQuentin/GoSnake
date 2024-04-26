package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth  = 320
	screenHeight = 240
	tileSize     = 5
)

type Game struct {
	snake    *Snake
	food     *Food
	renderer *Renderer
	logic    *GameLogic
}

type Point struct {
	X, Y int
}

type Snake struct {
	Body        []Point
	Direction   Point
	GrowCounter int
}

type Food struct {
	Position Point
}

type Renderer struct {
	screen *ebiten.Image
	face   font.Face
}

type GameLogic struct {
	score         int
	gameOver      bool
	gameWon       bool
	speed         int
	updateCounter int
}

func (g *Game) Update(screen *ebiten.Image) error {
	if g.logic.HandleGameState(inpututil.IsKeyJustPressed(ebiten.KeyR)) {
		g.restart()
		return nil
	}

	if g.logic.UpdateTick() {
		g.snake.updateDirection()
		g.logic.CheckCollisions(g.snake, g.food)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.screen = screen
	g.renderer.drawBackground()
	g.renderer.drawSnake(g.snake.Body)
	g.renderer.drawFood(g.food.Position)
	g.renderer.drawUI(g.logic.score, g.logic.gameOver, g.logic.gameWon)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) restart() {
	g.snake = NewSnake()
	g.food = NewFood()
	g.logic = NewGameLogic()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := &Game{
		snake:    NewSnake(),
		food:     NewFood(),
		renderer: NewRenderer(),
		logic:    NewGameLogic(),
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("GoSnake")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func NewSnake() *Snake {
	return &Snake{
		Body:      []Point{{X: screenWidth / tileSize / 2, Y: screenHeight / tileSize / 2}},
		Direction: Point{X: 1, Y: 0},
	}
}

func (s *Snake) updateDirection() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && s.Direction.X == 0 {
		s.Direction = Point{X: -1, Y: 0}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) && s.Direction.X == 0 {
		s.Direction = Point{X: 1, Y: 0}
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) && s.Direction.Y == 0 {
		s.Direction = Point{X: 0, Y: -1}
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) && s.Direction.Y == 0 {
		s.Direction = Point{X: 0, Y: 1}
	}
	s.Move()
}

func (s *Snake) Move() {
	newHead := Point{X: s.Body[0].X + s.Direction.X, Y: s.Body[0].Y + s.Direction.Y}
	s.Body = append([]Point{newHead}, s.Body...)
	if s.GrowCounter > 0 {
		s.GrowCounter--
	} else {
		s.Body = s.Body[:len(s.Body)-1]
	}
}

func NewFood() *Food {
	return &Food{
		Position: Point{X: rand.Intn(screenWidth / tileSize), Y: rand.Intn(screenHeight / tileSize)},
	}
}

func NewRenderer() *Renderer {
	return &Renderer{
		face: NewFaceWrapper(*basicfont.Face7x13),
	}
}

type FaceWrapper struct {
	basicfont.Face
}

func NewFaceWrapper(face basicfont.Face) font.Face {
	return &FaceWrapper{Face: face}
}

func (fw *FaceWrapper) Close() error {
	return nil
}

func (r *Renderer) drawBackground() {
	r.screen.Fill(color.RGBA{154, 198, 0, 255})
}

func (r *Renderer) drawSnake(body []Point) {
	for _, p := range body {
		ebitenutil.DrawRect(r.screen, float64(p.X*tileSize), float64(p.Y*tileSize), tileSize, tileSize, color.RGBA{33, 50, 15, 255}) // Snake color
	}
}

func (r *Renderer) drawFood(position Point) {
	ebitenutil.DrawRect(r.screen, float64(position.X*tileSize), float64(position.Y*tileSize), tileSize, tileSize, color.RGBA{231, 71, 29, 255}) // Food color
}

func (r *Renderer) drawUI(score int, gameOver bool, gameWon bool) {
	scoreText := fmt.Sprintf("Score: %d", score)
	text.Draw(r.screen, scoreText, r.face, 5, screenHeight-5, color.White)
	if gameOver {
		text.Draw(r.screen, "Game Over", r.face, screenWidth/2-40, screenHeight/2, color.White)
		text.Draw(r.screen, "Press 'R' to restart", r.face, screenWidth/2-60, screenHeight/2+16, color.White)
	}
	if gameWon {
		text.Draw(r.screen, "You Won!", r.face, screenWidth/2-40, screenHeight/2, color.White)
		text.Draw(r.screen, "Press 'R' to restart", r.face, screenWidth/2-60, screenHeight/2+16, color.White)
	}
}

func NewGameLogic() *GameLogic {
	return &GameLogic{speed: 10}
}

func (gl *GameLogic) HandleGameState(restartPressed bool) bool {
	if gl.gameOver || gl.gameWon {
		return restartPressed
	}
	return false
}

func (gl *GameLogic) UpdateTick() bool {
	gl.updateCounter++
	if gl.updateCounter < gl.speed {
		return false
	}
	gl.updateCounter = 0
	return true
}

func (gl *GameLogic) CheckCollisions(snake *Snake, food *Food) {
	head := snake.Body[0]
	if head.X < 0 || head.Y < 0 || head.X >= screenWidth/tileSize || head.Y >= screenHeight/tileSize {
		gl.gameOver = true
		gl.speed = 10
		return
	}

	for _, part := range snake.Body[1:] {
		if head.X == part.X && head.Y == part.Y {
			gl.gameOver = true
			gl.speed = 10
			return
		}
	}

	if head.X == food.Position.X && head.Y == food.Position.Y {
		gl.score++
		snake.GrowCounter += 1
		*food = *NewFood()

		if gl.score == 25 {
			gl.gameWon = true
			gl.speed = 10
		} else {
			if gl.speed > 2 {
				gl.speed--
			}
		}
	}
}

func (g *Game) restart() {
	g.snake = NewSnake()
	g.score = 0
	g.gameOver = false
	g.gameWon = false
	g.food = NewFood()
	g.speed = 10
}

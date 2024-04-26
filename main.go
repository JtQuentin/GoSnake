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

type GameEntity interface {
	Update()
	Draw(screen *ebiten.Image)
}

type EntityFactory interface {
	NewSnake() *Snake
	NewFood() *Food
	NewRenderer() *Renderer
	NewGameLogic() *GameLogic
}

type ConcreteFactory struct{}

func (f *ConcreteFactory) NewSnake() *Snake {
	return NewSnake()
}

func (f *ConcreteFactory) NewFood() *Food {
	return NewFood()
}

func (f *ConcreteFactory) NewRenderer() *Renderer {
	return NewRenderer()
}

func (f *ConcreteFactory) NewGameLogic() *GameLogic {
	return NewGameLogic()
}

type GameEntityController struct {
	snake *Snake
	food  *Food
}

func NewGameEntityController(factory EntityFactory) *GameEntityController {
	return &GameEntityController{
		snake: factory.NewSnake(),
		food:  factory.NewFood(),
	}
}

func (controller *GameEntityController) Update() {
	controller.snake.updateDirection()
	controller.food.Update()
}

func (controller *GameEntityController) Draw(screen *ebiten.Image) {
	controller.snake.Draw(screen)
	controller.food.Draw(screen)
}

type Game struct {
	entities *GameEntityController
	renderer *Renderer
	logic    *GameLogic
}

func NewGame(factory EntityFactory) *Game {
	renderer := factory.NewRenderer()
	logic := factory.NewGameLogic()
	entities := NewGameEntityController(factory)
	return &Game{
		entities: entities,
		renderer: renderer,
		logic:    logic,
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.restart()
	}
	if g.logic.UpdateTick() {
		g.entities.Update()
		g.logic.CheckCollisions(g.entities.snake, g.entities.food)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.Draw(screen, g.entities.snake.Body, g.entities.food.Position, g.logic.score, g.logic.gameOver, g.logic.gameWon)
}

func (g *Game) restart() {
	factory := &ConcreteFactory{}
	*g = *NewGame(factory)
}

type GameManager struct {
	game *Game
}

func NewGameManager(game *Game) *GameManager {
	return &GameManager{game: game}
}

func (gm *GameManager) Update(screen *ebiten.Image) error {
	return gm.game.Update()
}

func (gm *GameManager) Draw(screen *ebiten.Image) {
	gm.game.Draw(screen)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	factory := &ConcreteFactory{}
	game := NewGame(factory)
	gameManager := NewGameManager(game)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("GoSnake")
	if err := ebiten.RunGame(gameManager); err != nil {
		log.Fatal(err)
	}
}

type Point struct {
	X, Y int
}

type Snake struct {
	Body        []Point
	Direction   Point
	GrowCounter int
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

func (gm *GameManager) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (s *Snake) Draw(screen *ebiten.Image) {
	for _, p := range s.Body {
		ebitenutil.DrawRect(screen, float64(p.X*tileSize), float64(p.Y*tileSize), tileSize, tileSize, color.RGBA{33, 50, 15, 255})
	}
}

type Food struct {
	Position Point
}

func NewFood() *Food {
	return &Food{
		Position: Point{X: rand.Intn(screenWidth / tileSize), Y: rand.Intn(screenHeight / tileSize)},
	}
}

func (f *Food) Update() {
}

func (f *Food) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, float64(f.Position.X*tileSize), float64(f.Position.Y*tileSize), tileSize, tileSize, color.RGBA{231, 71, 29, 255})
}

type Renderer struct {
	screen *ebiten.Image
	face   font.Face
}

func NewRenderer() *Renderer {
	return &Renderer{
		face: basicfont.Face7x13,
	}
}

func (r *Renderer) Draw(screen *ebiten.Image, body []Point, position Point, score int, gameOver bool, gameWon bool) {
	r.screen = screen
	r.drawBackground()
	r.drawSnake(body)
	r.drawFood(position)
	r.drawUI(score, gameOver, gameWon)
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

type GameLogic struct {
	score         int
	gameOver      bool
	gameWon       bool
	speed         int
	updateCounter int
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

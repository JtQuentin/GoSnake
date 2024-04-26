package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"bufio"

	"os"
	"sort"
	"strings"

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
	snake        *Snake
	food         *Food
	renderer     *Renderer
	logic        *GameLogic
	startManager *GameStartManager
}

type Drawable interface {
	Draw(screen *ebiten.Image)
}

type Updatable interface {
	Update() error
}

type ScoreEntry struct {
	Name  string
	Score int
}

// SaveScore saves the current score to a file
func SaveScore(score int) error {
	file, err := os.OpenFile("scores.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Example: Save with a placeholder name and score
	_, err = file.WriteString(fmt.Sprintf("Player: %d\n", score))
	return err
}

// LoadScores loads scores from a file and returns a sorted slice of ScoreEntry
func LoadScores() ([]ScoreEntry, error) {
	file, err := os.Open("scores.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var scores []ScoreEntry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		if len(parts) == 2 {
			// Assuming the score is always an integer
			var score int
			fmt.Sscanf(parts[1], "%d", &score)
			scores = append(scores, ScoreEntry{Name: parts[0], Score: score})
		}
	}

	// Sort scores in descending order
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	return scores, scanner.Err()
}

type GameManager struct {
	game         *Game
	startManager *GameStartManager
	pauseManager *GamePauseManager
}

func NewGameManager(game *Game, startManager *GameStartManager, pauseManager *GamePauseManager) *GameManager {
	return &GameManager{game: game, startManager: startManager, pauseManager: pauseManager}
}

func (gm *GameManager) Update(screen *ebiten.Image) error {
	if !gm.startManager.IsGameStarted() {
		gm.startManager.HandleStartInput()
		return nil
	}

	if !gm.startManager.IsGameStarted() {
		gm.startManager.HandleStartInput()
		return nil
	}

	gm.pauseManager.HandlePauseInput()
	if gm.pauseManager.IsGamePaused() {
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		gm.game.restart()
	}

	if gm.game.logic.HandleGameState(inpututil.IsKeyJustPressed(ebiten.KeyR), gm.startManager.IsGameStarted()) {
		return nil
	}

	if gm.game.logic.UpdateTick() {
		gm.game.snake.updateDirection()
		gm.game.logic.CheckCollisions(gm.game.snake, gm.game.food)
	}

	gm.game.Draw(screen)

	return nil
}

func (gm *GameManager) Draw(screen *ebiten.Image) {
	gm.game.Draw(screen)
	gm.game.renderer.drawUI(gm.game.logic.score, gm.game.logic.gameOver, gm.game.logic.gameWon, gm.startManager.IsGameStarted())
}

func (gm *GameManager) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.screen = screen
	g.renderer.drawBackground()
	g.renderer.drawSnake(g.snake.Body)
	g.renderer.drawFood(g.food.Position)
	g.renderer.drawUI(g.logic.score, g.logic.gameOver, g.logic.gameWon, g.startManager.IsGameStarted())
}

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) restart() {
	g.snake = NewSnake()
	g.food = NewFood()
	g.logic = NewGameLogic()
}

func NewGame(snake *Snake, food *Food, renderer *Renderer, logic *GameLogic, startManager *GameStartManager) *Game {
	return &Game{
		snake:        snake,
		food:         food,
		renderer:     renderer,
		logic:        logic,
		startManager: startManager,
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	snake := NewSnake()
	food := NewFood()
	renderer := NewRenderer()
	logic := NewGameLogic()
	gameStartManager := NewGameStartManager()
	gamePauseManager := NewGamePauseManager()
	game := NewGame(snake, food, renderer, logic, gameStartManager)
	gameManager := NewGameManager(game, gameStartManager, gamePauseManager)

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
	if (ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft)) && s.Direction.X == 0 {
		s.Direction = Point{X: -1, Y: 0}
	} else if (ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight)) && s.Direction.X == 0 {
		s.Direction = Point{X: 1, Y: 0}
	} else if (ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp)) && s.Direction.Y == 0 {
		s.Direction = Point{X: 0, Y: -1}
	} else if (ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown)) && s.Direction.Y == 0 {
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

type GameStartManager struct {
	gameStart bool
}

func NewGameStartManager() *GameStartManager {
	return &GameStartManager{}
}

func (gsm *GameStartManager) HandleStartInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		gsm.gameStart = true
	}
}

func (gsm *GameStartManager) IsGameStarted() bool {
	return gsm.gameStart
}

func (gl *GameLogic) HandleGameState(restartPressed, gameStarted bool) bool {
	if !gameStarted || gl.gameOver || gl.gameWon {
		if restartPressed {
			gl.restartGame()
			return false
		}
		return true
	}
	return false
}

func (gl *GameLogic) restartGame() {
	gl.score = 0
	gl.gameOver = false
	gl.gameWon = false
	gl.speed = 10
	gl.updateCounter = 0
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
		SaveScore(gl.score)
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

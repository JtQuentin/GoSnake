package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"bufio"

	"os"
	"sort"
	"strings"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
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

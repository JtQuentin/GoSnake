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
)

const (
	screenWidth  = 320
	screenHeight = 240
	tileSize     = 5
)

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

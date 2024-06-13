package game

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// ScoreEntry represents a single entry in the score file
type ScoreEntry struct {
	Name  string // The name of the player
	Score int    // The score achieved by the player
}

// SaveScore saves the current score to a file
func SaveScore(score int) error {
	// If the score is 0, don't save it
	if score == 0 {
		return nil
	}

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
	// Open the score file
	file, err := os.Open("scores.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var scores []ScoreEntry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Read each line of the file
		line := scanner.Text()
		// Split the line into name and score
		parts := strings.Split(line, ": ")
		if len(parts) == 2 {
			// Parse the score as an integer
			var score int
			fmt.Sscanf(parts[1], "%d", &score)
			// Add the score entry to the list
			scores = append(scores, ScoreEntry{Name: parts[0], Score: score})
		}
	}

	// Sort the scores in descending order
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	// Return the scores and any error that occurred while reading the file
	return scores, scanner.Err()
}

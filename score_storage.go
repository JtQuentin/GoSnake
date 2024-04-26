package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

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

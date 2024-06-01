package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type GamePauseManager struct {
	gamePaused bool
}

func NewGamePauseManager() *GamePauseManager {
	return &GamePauseManager{}
}

func (gpm *GamePauseManager) HandlePauseInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		gpm.gamePaused = !gpm.gamePaused
	}
}

func (gpm *GamePauseManager) IsGamePaused() bool {
	return gpm.gamePaused
}

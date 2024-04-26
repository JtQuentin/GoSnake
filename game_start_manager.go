package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

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

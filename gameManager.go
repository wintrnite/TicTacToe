package main

import "ticTacToe/entities"

type GameManager struct {
	player entities.Player
	x      uint
	y      uint
}

func NewGameManager(player entities.Player, x, y uint) GameManager {
	return GameManager{player: player, x: x, y: y}
}

//func (gm GameManager) StartNewGame() error {
//	field := entities.NewField(gm.x, gm.y)
//	return nil
//}

func (gm *GameManager) playGame() error {
	for i := 0; i < int(gm.x*gm.y); i++ {

	}
	return nil
}

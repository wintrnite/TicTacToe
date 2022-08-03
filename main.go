package main

import (
	"log"
	"ticTacToe/entities"
	"ticTacToe/services"
)

func main() {
	x, y, err := services.ReadFieldSize()
	if err != nil {
		log.Fatal(err)
	}
	field := entities.NewField(x, y)
	playerMove, err := services.ReadMoveChoice()
	if err != nil {
		log.Fatal(err)
	}
	player := entities.NewPlayer(playerMove)
	gm := services.NewGameManager(player, field, x, y)
	err = gm.PlayGame()
	if err != nil {
		log.Fatal(err)
	}
}

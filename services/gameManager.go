package services

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"ticTacToe/entities"
)

type GameManager struct {
	player         entities.Player
	field          entities.Field
	x              uint
	y              uint
	botMove        entities.Move
	isPlayerTurn   bool
	lastPlacedCell entities.Cell
}

func NewGameManager(player entities.Player, field entities.Field, x, y uint) GameManager {
	gm := GameManager{player: player, field: field, x: x, y: y}
	if player.GetMoveChoice() == entities.MoveX {
		gm.botMove = entities.MoveO
		gm.isPlayerTurn = true
	} else {
		gm.botMove = entities.MoveX
	}
	return gm
}

func (gm *GameManager) PlayGame() error {
	for i := 0; i < int(gm.x*gm.y); i++ {
		if gm.isPlayerTurn {
			err := gm.makePlayerMove()
			if err != nil {
				return err
			}
			if gm.hasWin() {
				gm.endGame("вы победили")
				break
			}
		} else {
			gm.makeBotMove()
			if gm.hasWin() {
				gm.endGame("вы проиграли")
				break
			}
		}
		gm.isPlayerTurn = !gm.isPlayerTurn
	}
	return nil
}

func (gm *GameManager) makePlayerMove() (err error) {
	const maxIterCount = 5
	for i := 0; i < maxIterCount; i++ {
		fmt.Println(gm.field)
		x, y, err := ReadMove()
		if err != nil {
			return err
		}
		if !gm.field.InBounds(x, y) {
			err = fmt.Errorf("введенные координаты не влезают на поле")
			log.Println(err)
			continue
		}
		if gm.field.GetMove(x, y) != entities.EmptyMove {
			err = fmt.Errorf("сюда уже ходили")
			log.Println(err)
			continue
		} else {
			gm.field.SetMove(x, y, gm.player.GetMoveChoice())
			placedCell := entities.Cell{X: x, Y: y, Move: gm.player.GetMoveChoice()}
			gm.lastPlacedCell = placedCell
			break
		}
	}
	return err
}

func (gm *GameManager) makeBotMove() {
	var placed bool
	for !placed {
		x := uint(rand.Intn(int(gm.x)))
		y := uint(rand.Intn(int(gm.y)))
		if gm.field.GetMove(x, y) == entities.EmptyMove {
			gm.field.SetMove(x, y, gm.botMove)
			placedCell := entities.Cell{X: x, Y: y, Move: gm.botMove}
			gm.lastPlacedCell = placedCell
			placed = true
		}
	}

}

func (gm GameManager) endGame(endGameMessage string) {
	log.Println(endGameMessage)
}

func (gm GameManager) hasWin() bool {
	needToWin := int(math.Min(float64(gm.x), float64(gm.y))) - 1
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == j && j == 0 {
				continue
			}
			winCells1 := gm.GetWinMovesInDirection(i, j)
			winCells2 := gm.GetWinMovesInDirection(-i, -j)
			if len(winCells1)+len(winCells2) == needToWin {
				fmt.Println(gm.field)
				return true
			}
		}
	}
	return false
}

func (gm GameManager) GetWinMovesInDirection(i, j int) []entities.Cell {
	needToWin := int(math.Min(float64(gm.x), float64(gm.y)))
	winCells := make([]entities.Cell, 0)
	curX := uint(int(gm.lastPlacedCell.X) + i)
	curY := uint(int(gm.lastPlacedCell.Y) + j)
	neededMove := gm.lastPlacedCell.Move
	for len(winCells) != needToWin && gm.field.InBounds(curX, curY) && gm.field.GetMove(curX, curY) == neededMove {
		winCells = append(winCells, entities.Cell{X: curX, Y: curY, Move: neededMove})
		curX = uint(int(curX) + i)
		curY = uint(int(curY) + j)
	}
	return winCells
}

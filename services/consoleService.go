package services

import (
	"fmt"
	"log"
	"ticTacToe/entities"
)

const (
	maxIterCount = 5
)

var (
	valuesX = map[string]bool{"X": true, "x": true, "х": true, "Х": true}
	valuesO = map[string]bool{"O": true, "o": true, "О": true, "о": true, "0": true}
)

func ReadCoordinates() (uint, uint, error) {
	var x, y uint
	for i := 0; i < maxIterCount; i++ {
		fmt.Println("Введите коордитаты:\nВведите длину поля, нажмите Enter, затем введите ширину поля и нажмите Enter")
		_, err := fmt.Scan(&x, &y)
		if err != nil {
			log.Println("Не распознались координаты. Причина:", err)
			continue
		} else {
			break
		}
	}
	if x == 0 || y == 0 {
		return 0, 0, fmt.Errorf("некорректные координаты")
	}
	return x, y, nil
}

func ReadMoveChoice() (entities.Move, error) {
	var move string

	for i := 0; i < maxIterCount; i++ {
		fmt.Println("Введите сторону, за которую вы будете играть:\nX или O")
		_, err := fmt.Scan(&move)
		if err != nil {
			log.Println("Не распозналась сторона. Причина:", err)
			continue
		}
		_, ok1 := valuesX[move]
		_, ok2 := valuesO[move]
		if ok1 {
			return entities.MoveX, nil
		} else if ok2 {
			return entities.MoveO, nil
		}
	}
	return entities.EmptyMove, fmt.Errorf("не удалось распознать выбранную сторону")
}

//
//func ReadPlayerMoveCoordinates() (rune, rune, error) {
//	var x, y rune
//	var err error
//	for i := 0; i < maxIterCount; i++ {
//		err = nil
//		x, err = readCoordinate("Введите координату вашего хода по оси х (букву)")
//		if err != nil {
//			continue
//		}
//		y, err = readCoordinate("Введите координату вашего хода по оси y (число)")
//
//	}
//}

func readCoordinate(msg string) (string, error) {
	var x string
	fmt.Println(msg)
	_, err := fmt.Scan(&x)
	if err != nil {
		log.Println("Не распознались координату. Причина:", err)
		return 0, err
	}
	return x, nil
}

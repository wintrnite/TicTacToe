package services

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"ticTacToe/entities"
)

const (
	maxIterCount = 5
)

func ReadMoveChoice() (entities.Move, error) {
	var (
		move    string
		valuesX = map[string]bool{"X": true, "x": true, "х": true, "Х": true}
		valuesO = map[string]bool{"O": true, "o": true, "О": true, "о": true, "0": true}
	)
	for i := 0; i < maxIterCount; i++ {
		fmt.Println("Введите сторону, за которую вы будете играть:\nX или O")
		_, err := fmt.Scan(&move)
		if err != nil {
			log.Println("Не распозналась сторона. Причина:", err)
			continue
		}
		_, okX := valuesX[move]
		_, okO := valuesO[move]
		if okX {
			return entities.MoveX, nil
		} else if okO {
			return entities.MoveO, nil
		}
	}
	return entities.EmptyMove, fmt.Errorf("не удалось распознать выбранную сторону")
}

func ReadMove() (uint, uint, error) {
	xMessage := "Введите координату хода по оси Х (буква/строка, на латинице)"
	yMessage := "Введите координату хода по оси У (число)"
	x, y, err := readCoordinates(xMessage, yMessage, parseStrCoordinate, parseNumberCoordinate)
	if err != nil {
		return 0, 0, err
	}
	return x - 1, y - 1, nil
}

func ReadFieldSize() (uint, uint, error) {
	xMessage := "Введите длину поля"
	yMessage := "Введите высоту поля"
	return readCoordinates(xMessage, yMessage, parseNumberCoordinate, parseNumberCoordinate)
}

func readCoordinates(xMessage, yMessage string, parserX, parserY func(str string) (uint, error)) (uint, uint, error) {
	x, err := readCorrectCoordinateNTimes(xMessage, maxIterCount, parserX)
	if err != nil {
		return 0, 0, err
	}
	y, err := readCorrectCoordinateNTimes(yMessage, maxIterCount, parserY)
	if err != nil {
		return 0, 0, err
	}
	return x, y, nil
}

func readCorrectCoordinateNTimes(msg string, n uint, parser func(str string) (uint, error)) (uint, error) {
	var (
		coordinateStr string
		coordinate    uint
		err           error
		validateErr   error
	)
	for i := 0; i < int(n); i++ {
		coordinateStr, err = readCoordinateStr(msg)
		if err != nil {
			log.Println(err)
			continue
		}
		coordinate, validateErr = parser(coordinateStr)
		if validateErr != nil {
			err = validateErr
			log.Println(err)
			continue
		}
		if err == nil {
			break
		}
	}
	return coordinate, err
}

func parseNumberCoordinate(coordinate string) (uint, error) {
	coordinateInt, err := strconv.Atoi(coordinate)
	if err != nil {
		return 0, err
	}
	if coordinateInt < 0 {
		return 0, fmt.Errorf("координата должна быть больше 0")
	}
	return uint(coordinateInt), nil
}

func parseStrCoordinate(coordinate string) (uint, error) {
	var parsed uint
	runes := []rune(coordinate)
	runesLength := len(runes)
	getLetterValue := func(alphabetLetterOrder int, indexFromStart int) uint {
		return uint(math.Pow(27, float64(runesLength-indexFromStart-1))) * uint(alphabetLetterOrder)
	}
	for i := len(runes) - 1; i >= 0; i-- {
		v := runes[i]
		switch {
		case 65 <= v && v <= 90:
			alphabetLetterOrder := int(v - 64)
			parsed += getLetterValue(alphabetLetterOrder, i)
		case 97 <= v && v <= 122:
			alphabetLetterOrder := int(v - 96)
			parsed += getLetterValue(alphabetLetterOrder, i)
		default:
			return 0, fmt.Errorf(fmt.Sprintf("неккоректный сивмол в строке: %s", string(v)))
		}
	}
	return parsed, nil
}

func readCoordinateStr(msg string) (string, error) {
	var x string
	fmt.Println(msg)
	_, err := fmt.Scan(&x)
	if err != nil {
		return x, err
	}
	return x, nil
}

package main

import "fmt"

func main() {

	//x, y, err := services.ReadCoordinates()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(x, y)
	////moveChoice, err := services.ReadMoveChoice()
	////if err != nil {
	////	log.Fatal(err)
	////}
	//field := entities.NewField(x, y)
	//fmt.Println(field)
	//player := entities.NewPlayer(moveChoice)
	//gameManager := NewGameManager(player)
	//gameManager.StartNewGame()
	var a rune
	_, err := fmt.Scan(&a)
	if err != nil {
	}
	print(string(a))
}

package entities

type Player struct {
	move Move
}

func NewPlayer(moveChoice Move) Player {
	return Player{move: moveChoice}
}

func (p Player) GetMoveChoice() Move {
	return p.move
}

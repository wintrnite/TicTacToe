package entities

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

type Field struct {
	field [][]Move
	x     uint
	y     uint
}

func NewField(x, y uint) Field {
	matrix := make([][]Move, y)
	for i := range matrix {
		matrix[i] = make([]Move, x)
		for j := range matrix[i] {
			matrix[i][j] = EmptyMove
		}
	}
	return Field{field: matrix, x: x, y: y}
}

func (f *Field) SetMove(x, y uint, move Move) {
	f.field[y][x] = move
}

func (f Field) GetMove(x, y uint) Move {
	return f.field[y][x]
}

func (f Field) String() string {
	t := table.NewWriter()
	t.SetAutoIndex(true)
	for i := 0; i < len(f.field); i++ {
		t.AppendRow(f.getRow(i))
		t.AppendSeparator()
	}
	return t.Render()
}

func (f Field) InBounds(x, y uint) bool {
	return x < f.x && y < f.y
}

func (f Field) getRow(index int) table.Row {
	row := make([]interface{}, len(f.field[0]))
	for i, v := range f.field[index] {
		row[i] = string(v)
	}
	return row
}

package entities

import (
	"github.com/jedib0t/go-pretty/v6/table"
)

type Field struct {
	field [][]Move
}

func NewField(x, y uint) Field {
	matrix := make([][]Move, y)
	for i := range matrix {
		matrix[i] = make([]Move, x)
		for j := range matrix[i] {
			matrix[i][j] = EmptyMove
		}
	}
	return Field{field: matrix}
}

func (f *Field) SetMove(x, y uint, move Move) {
	f.field[x][y] = move
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

func (f Field) getRow(index int) table.Row {
	row := make([]interface{}, len(f.field[0]))
	for i, v := range f.field[index] {
		row[i] = string(v)
	}
	return row
}

func (f Field) check() {

}

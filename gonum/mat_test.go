package gonum

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/floats"

	"gonum.org/v1/gonum/mat"
)

func TestCol(t *testing.T) {
	m := mat.NewDense(3, 3, []float64{
		2.0, 9.0, 3.0,
		4.5, 6.7, 8.0,
		1.2, 3.0, 6.0,
	})

	row := mat.Row(nil, 1, m)
	col := mat.Col(nil, 1, m)
	maxValue := floats.Max(col)
	minValue := floats.Min(row)
	sub := maxValue - minValue
	sub = sub

	fmt.Printf("col = %#v", col)
}

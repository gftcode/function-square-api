package entity

import (
	"errors"
	"math"
)

type BaskaraVar struct {
	A int
	B int
	C int
	delta int
}

func NewEquation(a, b, c int) (*BaskaraVar, error) {
	if a == 0 {
		return nil, errors.New("o coeficiente 'A' não pode ser zero em uma equação quadrática")
	}
	return &BaskaraVar{A: a, B: b, C: c}, nil
}

func (e *BaskaraVar) FindDelta() int {
	e.delta = (e.B * e.B) - (4 * e.A * e.C) 
	return e.delta
}

func (e *BaskaraVar) HasXrows() bool {
	return e.delta >= 0
}

func (e *BaskaraVar) Xrows() (int, int) {
	raizDelta := math.Sqrt(float64(e.delta))

	x1 := (float64(-e.B) + raizDelta) / float64(2 * e.A)
	x2 := (float64(-e.B) - raizDelta) / float64(2 * e.A)

	return int(x1), int(x2)
}

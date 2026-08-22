package usecase

import (
	"fmt"
	"math"
)

type BaskaraVar struct {
	a     float64
	b     float64
	c     float64
	delta float64
}

type QuadraticFunction interface {
	FindDelta() float64
	Xrows() (float64, float64)
	FindVertices() (float64, float64)
}

func NewEquation(a, b, c int) (*BaskaraVar) {
	if a == 0 {
		fmt.Print("O coeficiente A não pode ser nulo..\nTrocando valor para '1'...")
		return &BaskaraVar{a: float64(1), b: float64(b), c: float64(c)}
	}

	return &BaskaraVar{a: float64(a), b: float64(b), c: float64(c)}
}

func (e *BaskaraVar) FindDelta() float64 {
	e.delta = (e.b * e.b) - (4 * e.a * e.c)
	return e.delta
}

func (e *BaskaraVar) Xrows() (float64, float64) {
	e.FindDelta()

	if e.delta < 0 {
		return 0, 0
	}

	raizDelta := math.Sqrt(e.delta)
	x1 := (-e.b + raizDelta) / (2 * e.a)
	x2 := (-e.b - raizDelta) / (2 * e.a)

	return x1, x2
}

func (e *BaskaraVar) FindVertices() (float64, float64) {
	e.FindDelta()
	
	Xv := -e.b / (2 * e.a)
	Yv := -e.delta / (4 * e.a)
	
	return Xv, Yv
}

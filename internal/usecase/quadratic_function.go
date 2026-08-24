package usecase

import (
	"fmt"
	"math"
)

type QuadraticFunction interface {
	GetCoefficients() (a, b, c float64)
	FindDelta() float64
	Xrows() (float64, float64)
	FindVertices() (float64, float64)
}

type Baskara struct {
	a     float64
	b     float64
	c     float64
	delta float64
}

func NewEquation(a, b, c int) QuadraticFunction {
	if a == 0 {
		a = 1
		fmt.Print("O coeficiente 'A' não pode ser '0'...\n")
		fmt.Print("Este coeficiente agora recebe o valor '1'...\n")
	}
	
	eq := &Baskara{a: float64(a), b: float64(b), c: float64(c)}
	eq.FindDelta()
	return eq
}

func (bask *Baskara) GetCoefficients() (a, b, c float64) {
	return bask.a, bask.b, bask.c
}

func (bask *Baskara) FindDelta() float64 {
	bask.delta = (bask.b * bask.b) - (4 * bask.a * bask.c)
	return bask.delta
}

func (bask *Baskara) Xrows() (float64, float64) {
	if bask.delta < 0 {
		return 0, 0
	}

	raizDelta := math.Sqrt(bask.delta)
	x1 := (-bask.b + raizDelta) / (2 * bask.a)
	x2 := (-bask.b - raizDelta) / (2 * bask.a)

	return x1, x2
}

func (bask *Baskara) FindVertices() (float64, float64) {
	Xv := -bask.b / (2 * bask.a)
	Yv := -bask.delta / (4 * bask.a)
	return Xv, Yv
}

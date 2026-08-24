package cli

import (
	"fmt"

	"github.com/gftcode/function-square-api/internal/usecase"
)

type Output struct {
	a     int
	b     int
	c     int
	delta float64
}

func NewCli() *Output {
	return &Output{}
}

func (o *Output) SetValues() usecase.QuadraticFunction {
	o.input()

	return usecase.NewEquation(o.a, o.b, o.c)
}

func (o *Output) GetValues() (int, int, int) {
	return o.a, o.b, o.c
}

func (o *Output) input() {
	var correct bool
	for {
		fmt.Print("Digite o valor de a (x²): ")
		fmt.Scan(&o.a)

		fmt.Print("Digite o valor de b (x): ")
		fmt.Scan(&o.b)

		fmt.Print("Digite o valor de c (termo constante): ")
		fmt.Scan(&o.c)

		fmt.Printf("\nSua equação: %dx² + (%dx) + (%d)\n", o.a, o.b, o.c)
		fmt.Print("Os dados estão corretos? (true/false): ")
		fmt.Scan(&correct)

		if correct {
			fmt.Println()
			break
		}
		fmt.Println("Vamos tentar de novo...")
	}
}
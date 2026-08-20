package main

import (
	"fmt"

	"github.com/gftcode/function-square-api/internal/usecase"
)

var x1, x2 int

func main() {
	function, err :=entity.NewEquation(1,10,4)
	delta := function.FindDelta()

	if err != nil{
		fmt.Printf("ERROR Equation: %s", err)
		return
	}

	if function.HasXrows() { 
		x1, x2 = function.Xrows()
		fmt.Printf("Raízes x1:%d | x2:%d\n", x1, x2)
		fmt.Printf("Delta: %d", delta)
	} else {
		fmt.Printf("Delta sem raíz: %d", delta)
	}
}
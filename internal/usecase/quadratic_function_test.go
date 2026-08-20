package entity_test

import (
	"fmt"
	"testing"

	"github.com/gftcode/function-square-api/internal/usecase"
)


func TestEquation(t *testing.T) {
	a, b, c := 1, 2, 3

	response ,_ := entity.NewEquation(a, b, c)
	delta := response.FindDelta()
	result := -8

	if delta != result {
		t.Fail()
		fmt.Printf("Valor esperado: %d.\nValor recebido: %d.\n", result, delta)
	}
}
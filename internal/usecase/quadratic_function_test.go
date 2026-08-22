package usecase_test

import (
	"testing"
	"github.com/gftcode/function-square-api/internal/usecase"
)

func TestDelta(t *testing.T) {
	a, b, c := 1, 2, 3
	expectedDelta := -8.0

	eq := usecase.NewEquation(a, b, c)
	delta := eq.FindDelta()

	if delta != expectedDelta {
		t.Errorf("Valor esperado: %f. Valor recebido: %f.", expectedDelta, delta)
	}
}

func TestXrows(t *testing.T) {
	a, b, c := -2, 8, -6
	expectedX1 := 1.0
	expectedX2 := 3.0

	eq := usecase.NewEquation(a, b, c)
	x1, x2 := eq.Xrows()

	if x1 != expectedX1 && x2 != expectedX2 {
		t.Errorf("Esperado x' e x'': {%f, %f}, Retornado x'e x'': {%f, %f}", expectedX1, expectedX2, x1, x2)
	}
}

func TestVertice(t *testing.T) {
	a, b, c := -2, 8, -6
	expectedXv := 2.0
	expectedYv := 2.0

	eq := usecase.NewEquation(a, b, c)
	Xv, Yv := eq.FindVertices()

	if Xv != expectedXv || Yv != expectedYv {
		t.Errorf("V esperado: {%f,%f} V retornado: {%f,%f}", expectedXv, expectedYv, Xv, Yv)
	}
}


package formula_test

import (
	"testing"

	"github.com/gftcode/function-square-api/internal/domain/formula"
)

func TestDeltaFormula(t *testing.T) {
	got := formula.NewFormula().DeltaFormula() 
	want := "Δ = b² - 4.a.c"
	if got != want {
		t.Errorf("Valor esperado: %q. Valor recebido: %q.", want, got)
	}
}

func TestXrowFormula(t *testing.T) {
	got := formula.NewFormula().XrowFormula()
	want := "x = (-b ± √Δ) / (2.a)"

	if got != want {
		t.Errorf("Valor esperado: %q. Valor recebido: %q.", want, got)
	} 
}

func TestVerticesFormula(t *testing.T) {
	got := formula.NewFormula().VerticesFormula()
	want := "Xv = -b / (2.a) | Yv = -Δ / (4.a)"

	if got != want {
		t.Errorf("Valor esperado: %q. Valor recebido: %q.", want, got)
	} 
}
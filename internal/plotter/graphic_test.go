package plotter

import "testing"

func TestNewTerminalPlotterNormalizesDimensions(t *testing.T) {
	plotter := NewTerminalPlotter(0, -2)
	
	if plotter.totalColumns != 1 || plotter.totalRows != 1 {
		t.Fatalf("dimensoes = (%d, %d), esperado (1, 1)", plotter.totalColumns, plotter.totalRows)
	}
}

func TestNewPlotBoundsIncludesRoots(t *testing.T) {
	bounds := newPlotBounds(2, 2, 1, 3, 16)

	xMinWant, xMaxWant := -4.0, 8.0
	yMinWant, yMaxWant := -3.0, 7.0

	if bounds.xMin != xMinWant || bounds.xMax != xMaxWant {
		t.Fatalf("limites X = (%v, %v), esperado (%v, %v)",
			bounds.xMin, bounds.xMax, xMinWant, xMaxWant)
	}

	if bounds.yMin != yMinWant || bounds.yMax != yMaxWant {
		t.Fatalf("limites Y = (%v, %v), esperado (%v, %v)",
			bounds.yMin, bounds.yMax, yMinWant, yMaxWant)
	}
}

func TestCanvasSetIgnoresOutOfBoundsPositions(t *testing.T) {
	canvas := newCanvas(2, 2, plotBounds{0, 1, 0, 1})
	canvas.set(-1, 0, "x")
	canvas.set(0, 2, "x")

	if canvas.cells[0][0] != " " || canvas.cells[1][1] != " " {
		t.Fatal("uma posicao invalida alterou o canvas")
	}
}

func TestRenderRowKeepsMultiCharacterMarkersAligned(t *testing.T) {
	got := renderRow([]string{"-", "X'", "-", "X''"})
	want := "-X'X''"

	if got != want {
		t.Fatalf("renderRow() = %q, esperado %q", got, want)
	}
}

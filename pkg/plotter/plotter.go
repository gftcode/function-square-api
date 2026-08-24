package plotter

import (
	"math"

	"github.com/gftcode/function-square-api/internal/usecase"
)

type TerminalPlotter struct {
	totalColumns int
	totalRows    int
}

func NewTerminalPlotter(columns, rows int) *TerminalPlotter {
	return &TerminalPlotter{totalColumns: maxInt(columns, 1), totalRows: maxInt(rows, 1)}
}

func (tp *TerminalPlotter) Plot(equation usecase.QuadraticFunction) {
	a, b, c := equation.GetCoefficients()
	delta := equation.FindDelta()
	xVertex, yVertex := equation.FindVertices()
	x1, x2 := equation.Xrows()

	bounds := newPlotBounds(xVertex, yVertex, x1, x2, delta)
	canvas := newCanvas(tp.totalColumns, tp.totalRows, bounds)
	canvas.drawAxes()
	canvas.drawCurve(a, b, c)
	canvas.mark(0, c, "C")
	if delta >= 0 {
		canvas.markRoots(x1, x2)
	}
	canvas.mark(xVertex, yVertex, "V")
	canvas.print()
}

type plotBounds struct {
	xMin, xMax float64
	yMin, yMax float64
}

func newPlotBounds(xVertex, yVertex, x1, x2, delta float64) plotBounds {
	bounds := plotBounds{xVertex - 5, xVertex + 5, yVertex - 5, yVertex + 5}
	if delta >= 0 {
		bounds.xMin = math.Min(bounds.xMin, math.Min(x1, x2)) - 1
		bounds.xMax = math.Max(bounds.xMax, math.Max(x1, x2)) + 1
	}
	return bounds
}

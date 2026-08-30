package plotter

import "math"

type canvas struct {
	columns, rows int
	bounds        plotBounds
	cells         [][]string
}

func newCanvas(columns, rows int, bounds plotBounds) *canvas {
	cells := make([][]string, rows)
	for row := range cells {
		cells[row] = make([]string, columns)
		for column := range cells[row] {
			cells[row][column] = " "
		}
	}
	return &canvas{columns: columns, rows: rows, bounds: bounds, cells: cells}
}

func (c *canvas) row(y float64) int {
	ratio := (c.bounds.yMax - y) / (c.bounds.yMax - c.bounds.yMin)
	return int(math.Round(ratio * float64(c.rows-1)))
}

func (c *canvas) column(x float64) int {
	ratio := (x - c.bounds.xMin) / (c.bounds.xMax - c.bounds.xMin)
	return int(math.Round(ratio * float64(c.columns-1)))
}

func (c *canvas) set(row, column int, value string) {
	if row >= 0 && row < c.rows && column >= 0 && column < c.columns {
		c.cells[row][column] = value
	}
}

func maxInt(first, second int) int {
	if first > second {
		return first
	}
	return second
}

package plotter

import (
	"math"
)

func (c *canvas) drawAxes() {
	for row := range c.cells {
		y := c.bounds.yMax - float64(row)/float64(maxInt(c.rows-1, 1))*(c.bounds.yMax-c.bounds.yMin)
		for column := range c.cells[row] {
			x := c.bounds.xMin + float64(column)/float64(maxInt(c.columns-1, 1))*(c.bounds.xMax-c.bounds.xMin)
			onY := math.Abs(y) < (c.bounds.yMax-c.bounds.yMin)/(float64(c.rows)*1.5)
			onX := math.Abs(x) < (c.bounds.xMax-c.bounds.xMin)/(float64(c.columns)*1.5)
			if onY {
				c.cells[row][column] = "-"
			}
			if onX {
				c.cells[row][column] = "|"
				if onY {
					c.cells[row][column] = "+"
				}
			}
		}
	}
}

func (c *canvas) drawCurve(a, b, constant float64) {
	for column := 0; column < c.columns; column++ {
		x := c.bounds.xMin + float64(column)/float64(maxInt(c.columns-1, 1))*(c.bounds.xMax-c.bounds.xMin)
		y := a*x*x + b*x + constant
		if y >= c.bounds.yMin && y <= c.bounds.yMax {
			c.set(c.row(y), column, "*")
		}
	}
}

func (c *canvas) mark(x, y float64, value string) {
	c.set(c.row(y), c.column(x), value)
}

func (c *canvas) markRoots(x1, x2 float64) {
	row := c.row(0)
	first, second := c.column(x1), c.column(x2)
	if first == second {
		c.set(row, first, "X")
		return
	}
	c.set(row, first, "X'")
	c.set(row, second, "X''")
}

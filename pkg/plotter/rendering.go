package plotter

import (
	"fmt"
	"strings"
)

func (c *canvas) print() {
	fmt.Println("\n=============================================")
	for _, row := range c.cells {
		fmt.Println(renderRow(row))
	}
	fmt.Println("=============================================")
}

func renderRow(row []string) string {
	var result strings.Builder
	skipNext := 0

	for i := 0; i < len(row); i++ {
		if skipNext > 0 {
			skipNext--
			continue
		}

		cell := row[i]
		result.WriteString(cell)
		if len(cell) > 1 {
			skipNext = len(cell) - 1
		}
	}

	return result.String()
}

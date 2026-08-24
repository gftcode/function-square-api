package main

import (
	"github.com/gftcode/function-square-api/internal/infrastructure/entrypoint/cli"
	"github.com/gftcode/function-square-api/pkg/plotter"
)

func main() {
	equation := cli.NewCli().SetValues()

	printer := plotter.NewTerminalPlotter(35, 12)

	printer.Plot(equation)
}

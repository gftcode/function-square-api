package main

import (
	"github.com/gftcode/function-square-api/internal/infrastructure/entrypoint/cli"
	"github.com/gftcode/function-square-api/internal/plotter"
)

func main() {
	equation := cli.NewCli().SetValues()

	printer := plotter.NewTerminalPlotter(35, 25)

	printer.Plot(equation)

}

package cli

import (
	"github.com/gftcode/function-square-api/internal/domain/formula"
	"github.com/gftcode/function-square-api/internal/usecase"
	"github.com/gftcode/function-square-api/pkg/fileutil"
)

func ExportResults(equation usecase.QuadraticFunction) error {
	file := fileutil.NewFile()
	prefix := ", "

	x, x2 := equation.Xrows()
	v, v2 := equation.FindVertices()

	if err := file.AddContent(formula.NewFormula().GetFormulas()); err != nil {
		return err
	}

	if err := file.AddContent("----------------------------------\n"); err != nil {
		return err
	}

	if err := file.AddContent("Delta: {", equation.FindDelta(), "}\n"); err != nil {
		return err
	}

	if x != 0 || x2 != 0 {
		if err := file.AddContent("X' and X'': {", x, prefix, x2, "}\n"); err != nil {
			return err
		}
	}

	if err := file.AddContent("Vertice: {", v, prefix, v2, "}\n"); err != nil {
		return err
	}

	return nil
}

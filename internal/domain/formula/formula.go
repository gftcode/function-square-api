package formula

type Formula struct {}

func NewFormula() *Formula {
	return &Formula{}
}

func (f *Formula) DeltaFormula() string {
	return "Δ = b² - 4.a.c"
}

func (f *Formula) XrowFormula() string {
	return "x = (-b ± √Δ) / (2.a)"
}

func (f *Formula) VerticesFormula() string {
	return "Xv = -b / (2.a) | Yv = -Δ / (4.a)"
}

func (f *Formula) GetFormulas() (string, string, string) {
	return f.DeltaFormula() + "\n", f.XrowFormula() + "\n", f.VerticesFormula()+ "\n"
}
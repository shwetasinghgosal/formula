package formula

import (
	"math"
)

type VariableDefine struct {
	A, B float64
}

func (f *VariableDefine) FormulaSquare() float64 {
	//(a+b)^2 = a^2+b^2+2ab
	return math.Pow(f.A, 2) + math.Pow(f.B, 2) + 2*f.A*f.B
}

package funcs

import (
	"github.com/Knetic/govaluate"
)

var functions = map[string]govaluate.ExpressionFunction{}

type Func interface {
	// Name func name
	Name() string

	// ExpressionFunction expression function description
	ExpressionFunction() govaluate.ExpressionFunction
}

func register(f Func) {
	functions[f.Name()] = f.ExpressionFunction()
}

func GetFunctions() map[string]govaluate.ExpressionFunction {
	return functions
}

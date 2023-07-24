package funcs

import (
	"fmt"
	"testing"

	"github.com/Knetic/govaluate"
)

func TestIF(t *testing.T) {
	fun := &IF{}

	fun.Name()

	functions := map[string]govaluate.ExpressionFunction{
		"IF": func(args ...interface{}) (interface{}, error) {
			condition, ok := args[0].(bool)
			if !ok {
				return nil, fmt.Errorf("invalid condition in IF function")
			}

			if condition {
				return args[1], nil
			} else {
				return args[2], nil
			}
		},
	}
	expression, err := govaluate.NewEvaluableExpressionWithFunctions("IF(1>2,1,2)", functions)
	if err != nil {
		return
	}

	evaluate, err := expression.Evaluate(nil)

	t.Log(evaluate)
}

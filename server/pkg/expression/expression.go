package expression

import (
	"github.com/Knetic/govaluate"

	"github.com/zhaoyunxing92/flexi-code/server/pkg/expression/funcs"
)

type EvaluableExpression struct {
	engine *govaluate.EvaluableExpression
}

func NewEvaluableExpression(expr string) *EvaluableExpression {

	if engine, err := govaluate.NewEvaluableExpressionWithFunctions(expr, funcs.GetFunctions()); err != nil {
		return nil
	} else {
		return &EvaluableExpression{engine: engine}
	}
}

func (e *EvaluableExpression) Evaluate(args map[string]interface{}) (interface{}, error) {
	return e.engine.Evaluate(args)
}

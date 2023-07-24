package funcs

import (
	"errors"

	"github.com/Knetic/govaluate"
)

func init() {
	register(&Sum{})
}

type Sum struct {
}

func (s *Sum) Name() string {
	return "SUM"
}

func (s *Sum) ExpressionFunction() govaluate.ExpressionFunction {
	return func(args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, errors.New("invalid number of parameters in「IF」function")
		}
		return args[0].(float64) + args[1].(float64), nil
	}
}

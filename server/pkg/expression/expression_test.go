package expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIFExpression(t *testing.T) {
	args := make(map[string]interface{})
	args["a"] = 5
	args["b"] = 8

	val, err := NewEvaluableExpression("SUM(IF(a, 1, 2),SUM(b,7))").Evaluate(args)
	assert.Nil(t, err)
	t.Log(val)
}

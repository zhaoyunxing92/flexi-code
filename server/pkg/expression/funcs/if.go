package funcs

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/Knetic/govaluate"
)

func init() {
	register(&IF{})
}

type IF struct {
}

func (i *IF) Name() string {
	return "IF"
}

func (i *IF) ExpressionFunction() govaluate.ExpressionFunction {
	return func(args ...interface{}) (interface{}, error) {
		var (
			condition bool
			err       error
		)
		if len(args) != 3 {
			return nil, errors.New("invalid number of parameters in「IF」function")
		}

		val := args[0]
		vt := reflect.TypeOf(val)
		// 判断具体类型
		switch vt.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println("Value is of type int")
		case reflect.String:
			condition, err = strconv.ParseBool(val.(string))
		default:
			fmt.Println("Unknown type")
		}

		if err != nil {
			return nil, errors.New("invalid number of parameters in「IF」function")
		}

		if condition {
			return args[1], nil
		} else {
			return args[2], nil
		}
	}
}

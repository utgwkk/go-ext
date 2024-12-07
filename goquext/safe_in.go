package goquext

import (
	"reflect"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

func safeIn(e exp.Inable, args ...any) exp.BooleanExpression {
	switch len(args) {
	case 0:
		return goqu.L("1").Eq(goqu.L("0"))
	case 1:
		arg := args[0]
		rv := reflect.ValueOf(arg)
		if rv.Kind() != reflect.Slice {
			return e.In(arg)
		}
		if rv.Len() == 0 {
			return goqu.L("1").Eq(goqu.L("0"))
		}
		return e.In(arg)
	default: // len(args) > 1
		// Fallback to len(args) == 1 case
		return e.In(args)
	}
}

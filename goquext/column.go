package goquext

import (
	"reflect"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type column struct {
	exp.IdentifierExpression
}

func C(col string) exp.IdentifierExpression {
	return column{goqu.C(col)}
}

func (c column) In(args ...any) exp.BooleanExpression {
	switch len(args) {
	case 0:
		return goqu.L("1").Eq(goqu.L("0"))
	case 1:
		arg := args[0]
		rv := reflect.ValueOf(arg)
		if rv.Kind() != reflect.Slice {
			return c.IdentifierExpression.In(arg)
		}
		if rv.Len() == 0 {
			return goqu.L("1").Eq(goqu.L("0"))
		}
		return c.IdentifierExpression.In(arg)
	default: // len(args) > 1
		// Fallback to len(args) == 1 case
		return c.In(args)
	}
}

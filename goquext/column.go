package goquext

import (
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
	return safeIn(c.IdentifierExpression, args...)
}

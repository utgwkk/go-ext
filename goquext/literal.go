package goquext

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type literal struct {
	exp.LiteralExpression
}

func L(lit string) exp.LiteralExpression {
	return literal{goqu.L(lit)}
}

func (l literal) In(args ...any) exp.BooleanExpression {
	return safeIn(l.LiteralExpression, args...)
}

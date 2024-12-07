package goqumysqlext

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

func DoUpdate(update interface{}) exp.ConflictUpdateExpression {
	return goqu.DoUpdate("", update)
}

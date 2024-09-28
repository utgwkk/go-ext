package mysqlsnippet

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

// ref: https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html

func asMySQLError(err error) (*mysql.MySQLError, bool) {
	mysqlErr := &mysql.MySQLError{}
	if !errors.As(err, &mysqlErr) {
		return nil, false
	}
	return mysqlErr, true
}

func IsErrDuplicateEntry(err error) bool {
	mysqlErr, ok := asMySQLError(err)
	if !ok {
		return false
	}
	return mysqlErr.Number == 1586
}

func IsErrDeadlockFound(err error) bool {
	mysqlErr, ok := asMySQLError(err)
	if !ok {
		return false
	}
	return mysqlErr.Number == 1213
}

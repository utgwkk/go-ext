package mysqldialect

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/dialect/mysql"
)

func SafeMySQLDialectOptions() *goqu.SQLDialectOptions {
	opts := mysql.DialectOptions()
	opts.SupportsInsertIgnoreSyntax = false
	return opts
}

func SafeMySQLDialectOptionsV8() *goqu.SQLDialectOptions {
	opts := mysql.DialectOptionsV8()
	opts.SupportsInsertIgnoreSyntax = false
	return opts
}

func init() {
	goqu.RegisterDialect("safemysql", SafeMySQLDialectOptions())
	goqu.RegisterDialect("safemysql8", SafeMySQLDialectOptionsV8())
}

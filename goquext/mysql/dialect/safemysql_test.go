package mysqldialect

import (
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

func ExampleSafeMySQLDialectOptionsV8() {
	type userRow struct {
		Id   string `db:"id"`
		Name string `db:"name"`
	}
	dialect := goqu.Dialect("safemysql8")

	var query string
	query, _, _ = dialect.Insert("users").Rows(userRow{Id: "1", Name: "John"}).ToSQL()
	fmt.Println("normal update: " + query)

	query, _, _ = dialect.Insert("users").Rows(userRow{Id: "1", Name: "John"}).As("new").
		OnConflict(goqu.DoUpdate("", goqu.Record{"name": goqu.T("new").Col("name")})).ToSQL()
	fmt.Println("upsert: " + query)

	// Output:
	// normal update: INSERT INTO `users` (`id`, `name`) VALUES ('1', 'John')
	// upsert: INSERT INTO `users` (`id`, `name`) VALUES ('1', 'John') AS `new` ON DUPLICATE KEY UPDATE `name`=`new`.`name`
}

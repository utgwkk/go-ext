package goquext

import (
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
)

func ExampleColumnIn() {
	dialect := goqu.Dialect("sqlite3")
	var query string

	query, _, _ = dialect.From("users").Where(C("id").In([]int{})).ToSQL()
	fmt.Println("empty slice: " + query)

	query, _, _ = dialect.From("users").Where(C("id").In([]int{1, 2})).ToSQL()
	fmt.Println("non empty slice: " + query)

	query, _, _ = dialect.From("users").Where(C("id").In(1)).ToSQL()
	fmt.Println("single arg: " + query)

	query, _, _ = dialect.From("users").Where(C("id").In(1, 2)).ToSQL()
	fmt.Println("variadic args: " + query)

	query, _, _ = dialect.From("users").Where(C("id").In()).ToSQL()
	fmt.Println("empty variadic arg: " + query)

	// Output:
	// empty slice: SELECT * FROM `users` WHERE (1 = 0)
	// non empty slice: SELECT * FROM `users` WHERE (`id` IN (1, 2))
	// single arg: SELECT * FROM `users` WHERE (`id` IN (1))
	// variadic args: SELECT * FROM `users` WHERE (`id` IN (1, 2))
	// empty variadic arg: SELECT * FROM `users` WHERE (1 = 0)
}

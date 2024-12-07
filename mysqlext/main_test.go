package mysqlext

import (
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest/v3"
)

var db *sqlx.DB

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("failed to construct pool: %s", err)
	}

	if err := pool.Client.Ping(); err != nil {
		log.Fatalf("failed to ping: %s", err)
	}

	resource, err := pool.Run("mysql", "8", []string{"MYSQL_ROOT_PASSWORD=password"})
	if err != nil {
		log.Fatalf("failed to run container: %s", err)
	}

	if err := pool.Retry(func() error {
		dsn := fmt.Sprintf("root:password@tcp(localhost:%s)/mysql", resource.GetPort("3306/tcp"))
		dbConn, err := sqlx.Open("mysql", dsn)
		if err != nil {
			return err
		}
		db = dbConn
		return dbConn.Ping()
	}); err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}
	defer func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("failed to purge resource: %s", err)
		}
	}()

	exitcode := m.Run()
	os.Exit(exitcode)
}

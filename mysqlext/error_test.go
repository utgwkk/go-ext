package mysqlext

import "testing"

func TestIsErrDuplicateEntry(t *testing.T) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS test_table (
			id INT PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(255) UNIQUE
		)
	`)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO test_table (name) VALUES ('test')`)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO test_table (name) VALUES ('test')`)
	if err == nil {
		t.Fatal("want error, got nil")
	}

	if !IsErrDuplicateEntry(err) {
		t.Error("IsErrDuplicateEntry(err): want true, got false")
	}
}

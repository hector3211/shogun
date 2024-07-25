package shogun

import "testing"

func TestCt(t *testing.T) {
	query := "CREATE TABLE IF NOT EXISTS users (id INT NOT NULL PRIMARY KEY,name TEXT NOT NULL);"
	ct := NewCreateTableBuilder().
		CreaetTable("users").
		IfNotExists().
		Define("id", "INT", "NOT NULL", "PRIMARY KEY").
		Define("name", "TEXT", "NOT NULL").
		String()

	if ct != query {
		t.Fatalf("failed! testCt wanted: %s, got %s", query, ct)
	}
}

func TestCtTwo(t *testing.T) {
	query := "CREATE TABLE IF NOT EXISTS users (id INT NOT NULL PRIMARY KEY);"
	ct := NewCreateTableBuilder().
		CreaetTable("users").
		IfNotExists().
		Define("id", "INT", "NOT NULL", "PRIMARY KEY").
		String()

	if ct != query {
		t.Fatalf("failed! testCtTwo wanted: %s, got %s", query, ct)
	}
}

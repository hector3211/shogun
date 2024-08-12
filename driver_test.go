package shogun

import (
	"testing"
)

func TestDriverString(t *testing.T) {
	stmt := SQLITE.NewSelectBuilder().Select("*").From("users").OrderBy("name").Asc().Limit(15)
	driver := stmt.GetDriver()

	if driver != "sqlite3" {
		t.Fatalf("TestDriverString failed wanted 'sqlite3' got %s", driver)
	}
}

func TestDriverSqlite(t *testing.T) {
	query := "SELECT * FROM users ORDER BY name ASC LIMIT 15;"

	stmt := SQLITE.NewSelectBuilder().Select("*").From("users").OrderBy("name").Asc().Limit(15).Build()

	if stmt != query {
		t.Fatalf("TestDriverSqlite failed wanted %s got %s", query, stmt)
	}
}

func TestDriverPostgres(t *testing.T) {
	query := "SELECT * FROM users ORDER BY name ASC LIMIT 15;"

	stmt := POSTGRES.NewSelectBuilder().Select("*").From("users").OrderBy("name").Asc().Limit(15).Build()

	if stmt != query {
		t.Fatalf("TestDriverPostgres failed wanted %s got %s", query, stmt)
	}
}

func TestDriverInsert(t *testing.T) {
	query := "INSERT INTO users (name) VALUES ('maddog');"

	insertQuery := POSTGRES.NewInsertBuilder().Insert("users").Cols("name").Vals("maddog").Build()

	if insertQuery != query {
		t.Fatalf("TestDriverInsert failed, wanted %s got %s", query, insertQuery)
	}
}

func TestDriverUpdate(t *testing.T) {
	query := "UPDATE users SET name = 'maddog' WHERE name = 'hector';"

	stmt := SQLITE.NewUpdateBuilder().
		Update("users").
		Set(Equal("name", "maddog")).
		Where(Equal("name", "hector")).
		Build()

	if stmt != query {
		t.Fatalf("TestDriverUpdate failed wanted %s got %s", query, stmt)
	}
}

func TestDriverDelete(t *testing.T) {
	query := "DELETE users WHERE id = 1 OR name = 'hector' AND verifyEmail = FALSE;"

	stmt := POSTGRES.NewDeleteBuilder().Delete("users").Where(
		Equal("id", 1),
		Or(),
		Equal("name", "hector"),
		And(),
		Equal("verifyEmail", false)).Build()

	if stmt != query {
		t.Fatalf("TestDriverUpdate failed wanted %s got %s", query, stmt)
	}
}

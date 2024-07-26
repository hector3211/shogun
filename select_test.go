package shogun

import "testing"

func TestSelectOne(t *testing.T) {
	query := "SELECT * FROM users;"

	stmt := NewCreateSelectBuilder().Select("*").From("users").Build()
	// t.Fatalf("length: %d", len(stmt.Args))

	if stmt != query {
		t.Fatalf("TestSelectOne failed wanted %s got %s", query, stmt)
	}
}

func TestSelect(t *testing.T) {
	query := "SELECT (id) FROM users;"
	stmt := NewCreateSelectBuilder().Select("id").From("users").Build()

	if stmt != query {
		t.Fatalf("TestSelectTwo failed wanted %s got %s", query, stmt)
	}
}

func TestSelectMultipleFields(t *testing.T) {
	query := "SELECT (id,name) FROM users;"
	stmt := NewCreateSelectBuilder().Select("id,name").From("users").Build()

	if stmt != query {
		t.Fatalf("TestSelectTwo failed wanted %s got %s", query, stmt)
	}
}

func TestSelectMultipleTables(t *testing.T) {
	query := "SELECT (id,name) FROM (users,products);"
	stmt := NewCreateSelectBuilder().Select("id", "name").From("users", "products").Build()

	if stmt != query {
		t.Fatalf("TestSelectMultipleTables failed wanted %s got %s", query, stmt)
	}
}

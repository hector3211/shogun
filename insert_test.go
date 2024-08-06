package shogun

import "testing"

func TestInsertSimple(t *testing.T) {
	query := "INSERT INTO users (name) VALUES ('maddog');"

	insertQuery := NewInsertBuilder().Table("users").Cols("name").Vals("maddog").Build()

	if insertQuery != query {
		t.Fatalf("TestinsertSimple failed, wanted %s got %s", query, insertQuery)
	}
}

func TestInsertDouble(t *testing.T) {
	query := "INSERT INTO users (name,age) VALUES ('maddog',20);"

	insertQuery := NewInsertBuilder().Table("users").Cols("name", "age").Vals("maddog", 20).Build()

	if insertQuery != query {
		t.Fatalf("TestinsertDouble failed, wanted %s got %s", query, insertQuery)
	}
}

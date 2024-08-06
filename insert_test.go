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

func TestInsertBool(t *testing.T) {
	query := "INSERT INTO users (name,age,email_verified) VALUES ('maddog',20,TRUE);"

	insertQuery := NewInsertBuilder().Table("users").Cols("name", "age", "email_verified").Vals("maddog", 20, true).Build()

	if insertQuery != query {
		t.Fatalf("TestinsertDouble failed, wanted %s got %s", query, insertQuery)
	}
}

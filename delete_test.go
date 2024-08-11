package shogun

import "testing"

func TestDelete(t *testing.T) {
	query := "DELETE users WHERE id = 1 AND name = 'hector';"

	stmt := Delete("users").Where(Equal("id", 1), And(), Equal("name", "hector")).Build()

	if stmt != query {
		t.Fatalf("TestDelete failed! wanted %s got %s", query, stmt)
	}
}

package shogun

import "testing"

func TestDelete(t *testing.T) {
	query := "DELETE users WHERE id = 1 AND name = 'hector';"

	stmt := Delete("users").Where(Equal("id", 1), And(), Equal("name", "hector")).Build()

	if stmt != query {
		t.Fatalf("TestDelete failed! wanted %s got %s", query, stmt)
	}
}

func TestDeleteTwo(t *testing.T) {
	query := "DELETE users WHERE id = 1 AND name = 'hector' AND verifyEmail = TRUE;"

	stmt := Delete("users").Where(
		Equal("id", 1),
		And(),
		Equal("name", "hector"),
		And(),
		Equal("verifyEmail", true)).Build()

	if stmt != query {
		t.Fatalf("TestDeleteTwo failed! wanted %s got %s", query, stmt)
	}
}

func TestDeleteThree(t *testing.T) {
	query := "DELETE users WHERE id = 1 OR name = 'hector' AND verifyEmail = FALSE;"

	stmt := NewDeleteBuilder().Delete("users").Where(
		Equal("id", 1),
		Or(),
		Equal("name", "hector"),
		And(),
		Equal("verifyEmail", false)).Build()

	if stmt != query {
		t.Fatalf("TestDeleteThree failed! wanted %s got %s", query, stmt)
	}
}

package shogun

import "testing"

func TestUpdateOne(t *testing.T) {
	query := "UPDATE users SET name = 'maddog' WHERE name = 'hector';"

	stmt := NewUpdateBuilder().
		Update("users").
		Set(Equal("name", "maddog")).
		Where(Equal("name", "hector")).
		Build()

	if stmt != query {
		t.Fatalf("TestUpdateOne failed wanted %s got %s", query, stmt)
	}

}

func TestUpdateTwo(t *testing.T) {
	query := "UPDATE users SET name = 'maddog' WHERE name = 'hector' AND id = 1;"

	stmt := NewUpdateBuilder().
		Update("users").
		Set(Equal("name", "maddog")).
		Where(Equal("name", "hector"), And(), Equal("id", 1)).
		Build()

	if stmt != query {
		t.Fatalf("TestUpdateTwo failed wanted %s got %s", query, stmt)
	}

}

package shogun

import "testing"

func TestSelectOne(t *testing.T) {
	query := "SELECT * FROM users;"

	stmt := NewSelectBuilder().Select("*").From("users").Build()
	// t.Fatalf("length: %d", len(stmt.Args))

	if stmt != query {
		t.Fatalf("TestSelectOne failed wanted %s got %s", query, stmt)
	}
}

func TestSelectTwo(t *testing.T) {
	query := "SELECT (id) FROM users;"
	stmt := NewSelectBuilder().Select("id").From("users").Build()

	if stmt != query {
		t.Fatalf("TestSelectTwo failed wanted %s got %s", query, stmt)
	}
}

func TestSelectThree(t *testing.T) {
	query := "SELECT (id,name) FROM users;"
	stmt := NewSelectBuilder().Select("id,name").From("users").Build()

	if stmt != query {
		t.Fatalf("TestSelectThree failed wanted %s got %s", query, stmt)
	}
}

func TestSelectFour(t *testing.T) {
	query := "SELECT (id,name) FROM (users,products);"
	stmt := NewSelectBuilder().Select("id", "name").From("users", "products").Build()

	if stmt != query {
		t.Fatalf("TestSelectFour failed wanted %s got %s", query, stmt)
	}
}

func TestSelectFive(t *testing.T) {
	query := `SELECT (id,name) FROM users WHERE name = 'hector';`
	stmt := NewSelectBuilder().
		Select("id", "name").
		From("users").
		Where(
			Equal("name", "hector"),
		).
		Build()

	if stmt != query {
		t.Fatalf("TestSelectFive failed wanted %s got %s", query, stmt)
	}
}

func TestSelectWitNotEqual(t *testing.T) {
	query := `SELECT (id,name) FROM users WHERE name != 'hector';`
	stmt := NewSelectBuilder().
		Select("id", "name").
		From("users").
		Where(
			NotEqual("name", "hector"),
		).
		Build()

	if stmt != query {
		t.Fatalf("TestSelectWitNotEqual failed wanted %s got %s", query, stmt)
	}
}

func TestSelectWithAnd(t *testing.T) {
	query := `SELECT (id,name) FROM users WHERE name = 'hector' AND id < 10;`
	stmt := NewSelectBuilder().
		Select("id", "name").
		From("users").
		Where(
			Equal("name", "hector"),
			And(),
			LessThan("id", 10),
		).
		Build()

	if stmt != query {
		t.Fatalf("TestSelectWithAnd failed wanted %s got %s", query, stmt)
	}
}

func TestSelectWithOr(t *testing.T) {
	query := `SELECT (id,name) FROM users WHERE name = 'hector' OR id > 10;`
	stmt := NewSelectBuilder().
		Select("id", "name").
		From("users").
		Where(
			Equal("name", "hector"),
			Or(),
			GreaterThan("id", 10),
		).
		Build()

	if stmt != query {
		t.Fatalf("TestSelectWithOr failed wanted %s got %s", query, stmt)
	}
}

// TODO: run a test with bool values

package shogun

import "testing"

func TestSelect(t *testing.T) {
	query := "SELECT * FROM users;"

	stmt := Select("*").From("users").Build()

	if stmt != query {
		t.Fatalf("TestSelectOne failed wanted %s got %s", query, stmt)
	}
}

func TestSelectLimit(t *testing.T) {
	query := "SELECT * FROM users LIMIT 5;"

	stmt := Select("*").From("users").Limit(5).Build()

	if stmt != query {
		t.Fatalf("TestSelectLimit failed wanted %s got %s", query, stmt)
	}
}

func TestSelectOne(t *testing.T) {
	query := "SELECT * FROM users;"

	stmt := NewSelectBuilder().Select("*").From("users").Build()

	if stmt != query {
		t.Fatalf("TestSelectOne failed wanted %s got %s", query, stmt)
	}
}

func TestSelectTwo(t *testing.T) {
	query := "SELECT id FROM users;"
	stmt := NewSelectBuilder().Select("id").From("users").Build()

	if stmt != query {
		t.Fatalf("TestSelectTwo failed wanted %s got %s", query, stmt)
	}
}

func TestSelectThree(t *testing.T) {
	query := "SELECT id,name FROM users;"
	stmt := NewSelectBuilder().Select("id,name").From("users").Build()

	if stmt != query {
		t.Fatalf("TestSelectThree failed wanted %s got %s", query, stmt)
	}
}

func TestSelectFour(t *testing.T) {
	query := "SELECT id,name FROM users,products;"
	stmt := NewSelectBuilder().Select("id", "name").From("users", "products").Build()

	if stmt != query {
		t.Fatalf("TestSelectFour failed wanted %s got %s", query, stmt)
	}
}

func TestSelectFive(t *testing.T) {
	query := `SELECT id,name FROM users WHERE name = 'hector';`
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
	query := `SELECT id,name FROM users WHERE name != 'hector';`
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
	query := `SELECT id,name FROM users WHERE name = 'hector' AND id < 10;`
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
	query := `SELECT id,name FROM users WHERE name = 'hector' OR id > 10;`
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

func TestSelectOrder(t *testing.T) {
	query := "SELECT * FROM users ORDER BY name DESC LIMIT 5;"

	stmt := Select("*").From("users").OrderBy("name").Desc().Limit(5).Build()
	// t.Fatalf("length: %d", len(stmt.Args))

	if stmt != query {
		t.Fatalf("TestSelectOrder failed wanted %s got %s", query, stmt)
	}
}

func TestSelectOrderTwo(t *testing.T) {
	query := "SELECT * FROM users ORDER BY name ASC LIMIT 15;"

	stmt := Select("*").From("users").OrderBy("name").Asc().Limit(15).Build()
	// t.Fatalf("length: %d", len(stmt.Args))

	if stmt != query {
		t.Fatalf("TestSelectOrder failed wanted %s got %s", query, stmt)
	}
}

func TestSelectGroubyAndHaving(t *testing.T) {
	query := "SELECT sales_person_id,product_id FROM sales GROUP BY sales_person_id,product_id HAVING COUNT(*) > 10;"

	stmt := Select("sales_person_id", "product_id").From("sales").GroupBy("sales_person_id", "product_id").Having(Count("*", GREATERTHAN, 10))

	if stmt.Build() != query {
		t.Fatalf("TestSelectGroubyAndHaving failed! wanted %s got %s", query, stmt.Build())
	}

}

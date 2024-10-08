package shogun

import "testing"

func TestInsert(t *testing.T) {
	query := "INSERT INTO users (name) VALUES ('maddog');"

	insertQuery := Insert("users").Columns("name").Values("maddog").Build()

	if insertQuery != query {
		t.Fatalf("TestinsertSimple failed, wanted %s got %s", query, insertQuery)
	}
}

func TestInsertSimple(t *testing.T) {
	query := "INSERT INTO users (name) VALUES ('maddog');"

	insertQuery := NewInsertBuilder().Insert("users").Columns("name").Values("maddog").Build()

	if insertQuery != query {
		t.Fatalf("TestinsertSimple failed, wanted %s got %s", query, insertQuery)
	}
}

func TestInsertDouble(t *testing.T) {
	query := "INSERT INTO users (name,age) VALUES ('maddog',20);"

	insertQuery := NewInsertBuilder().Insert("users").Columns("name", "age").Values("maddog", 20).Build()

	if insertQuery != query {
		t.Fatalf("TestinsertDouble failed, wanted %s got %s", query, insertQuery)
	}
}

func TestInsertUpsert(t *testing.T) {
	query := "INSERT INTO users (id,name) VALUES (1,'Alice') ON CONFLICT(id) DO UPDATE SET name = 'NewAlice';"

	insertQuery := Insert("users").Columns("id", "name").Values(1, "Alice").OnConflict("id").DoUpdate("name", "NewAlice").Build()

	if insertQuery != query {
		t.Fatalf("TestInsertUpsert failed, wanted %s got %s", query, insertQuery)
	}
}

func TestInsertUpsertTwo(t *testing.T) {
	query := "INSERT INTO users (name,email) VALUES ('doe','email@email.com') ON CONFLICT(email) DO UPDATE SET email = 'Newemail@email.com';"

	insertQuery := Insert("users").Columns("name", "email").Values("doe", "email@email.com").OnConflict("email").DoUpdate("email", "Newemail@email.com")

	if insertQuery.Build() != query {
		t.Fatalf("TestInsertUpsert failed, wanted %s got %s", query, insertQuery)
	}
}

func TestInsertUpsertThree(t *testing.T) {
	query := "INSERT INTO users (email,verified_email) VALUES ('email@email.com',TRUE) ON CONFLICT(email) DO UPDATE SET verified_email = TRUE;"

	insertQuery := Insert("users").Columns("email", "verified_email").Values("email@email.com", true).OnConflict("email").DoUpdate("verified_email", true).Build()

	if insertQuery != query {
		t.Fatalf("TestInsertUpsert failed, wanted %s got %s", query, insertQuery)
	}
}

package shogun

import "testing"

func TestUpsert(t *testing.T) {
	query := "INSERT INTO users (id,name) VALUES (1,'Alice') ON CONFLICT(id) DO UPDATE SET name = 'NewAlice';"

	stmt := Upsert("users").Cols("id", "name").Vals(1, "Alice").ConflictOn("id").Update("name", "NewAlice").Build()
	if stmt != query {
		t.Fatalf("TestUpsert failed! wanted %s got %s", query, stmt)
	}
}

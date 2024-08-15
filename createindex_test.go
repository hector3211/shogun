package shogun

import "testing"

func TestIndex(t *testing.T) {
	query := "CREATE INDEX idx_name ON users(name);"

	stmt := Index("idx_name").On("users", "name").Build()
	if stmt != query {
		t.Fatalf("TestIndex failed! wanted %s got %s", query, stmt)
	}
}

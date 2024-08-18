package shogun

import "testing"

func TestTransaction(t *testing.T) {
	query := "BEGIN TRANSACTION; UPDATE users SET name = 'maddog' WHERE name = 'hector'; INSERT INTO users (name) VALUES ('maddog'); COMMIT; ROLLBACK;"

	update := Update("users").
		Set(Equal("name", "maddog")).
		Where(Equal("name", "hector"))

	insert := Insert("users").Columns("name").Values("maddog")

	transaction := UpdateTransaction(update).InsertTransaction(insert).Commit().RollBack()

	if transaction.Build() != query {
		t.Fatalf("TestTransaction failed! wanted %s got %s", query, transaction.Build())
	}
}

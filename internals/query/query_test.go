package query

import (
	"fmt"
	"shogun/utils"
	"testing"
)

type UsersTest struct {
	ID   int    `orm:"id"`
	Name string `orm:"name"`
}

func TestCreatingTable(t *testing.T) {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS userstest (id %s,name %s);", utils.PRIMARYKEY, utils.TEXT)

	userOne := UsersTest{
		ID:   0,
		Name: "hector",
	}

	resultQuery := GenerateNewTable(userOne, utils.Postgres)

	if resultQuery != query {
		t.Fatalf("failed creating new table!")
	}
}

func TestSelectStatement(t *testing.T) {
	fields := []string{"id", "name"}
	var query string
	if len(fields) > 0 {
		query = "SELECT (id,name) FROM userstest;"
	} else {
		query = "SELECT * FROM userstest;"
	}

	resultQuery := GenerateNewSelectStatement("userstest", "id,name")

	if resultQuery != query {
		t.Fatalf("failed! Select querys dont match up!")
	}
}

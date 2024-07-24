package orm

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"ormer/internals/query"
	"ormer/utils"
	"reflect"
	"strings"
)

// type Table interface {
// 	TableName() string
// }

// ORM object.
type Orm struct {
	DB     *sql.DB
	Tables []string
	Driver utils.Driver
}

// Creates a new instance of the ORM object.
func NewOrm(databaseDriver utils.Driver, connectionString string) (*Orm, error) {
	driver := utils.MatchDriver(databaseDriver)
	db, err := sql.Open(driver, connectionString)
	if err != nil {
		return &Orm{}, fmt.Errorf("failed creating new ORM: %s", err.Error())
	}

	return &Orm{DB: db, Driver: databaseDriver}, nil
}

// Creates a new table based on a struct given.
func (o *Orm) CreateNewTable(table interface{}) error {
	t := reflect.TypeOf(table)

	if t.Kind() != reflect.Struct {
		return fmt.Errorf("table struct passed is invalid")
	}
	query := query.GenerateNewTable(table, o.Driver)
	_, err := o.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed creating new table: %s", err.Error())
	}

	tableName := strings.ToLower(t.Name())
	o.Tables = append(o.Tables, tableName)

	return nil
}

// NewSelectStatement genereate a new select query statement based on the table passed and fields to query.
func (o *Orm) NewSelectStatement(table interface{}, fieldsSelectd []string) (*sql.Rows, error) {
	query := query.GenerateNewSelectStatement(table, fieldsSelectd) // need the struct here
	result, err := o.DB.Query(query)
	// defer result.Close()
	if err != nil {
		return nil, fmt.Errorf("failed executing select satement: %s", err.Error())
	}

	return result, nil
}

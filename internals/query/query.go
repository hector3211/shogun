package query

import (
	"fmt"
	"ormer/utils"
	"reflect"
	"strings"
)

// Generates a new table creation statement based on a struct and database driver.
func GenerateNewTable(table interface{}, driver utils.Driver) string {
	t := reflect.TypeOf(table)

	if t.Kind() != reflect.Struct {
		return ""
	}

	tableName := strings.ToLower(t.Name())
	var columns []string
	// fmt.Printf("Table name: %s\n", tableName)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		ormTag := field.Tag.Get("orm")
		sqlType := utils.ToSqlTypes(field.Type.Name())

		if ormTag == "id" {
			columns = append(columns, fmt.Sprintf("%s %s", ormTag, utils.PRIMARYKEY))
		} else {
			columns = append(columns, fmt.Sprintf("%s %s", ormTag, sqlType))
		}
	}

	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", tableName, strings.Join(columns, ","))
}

// Generates a new SELECT statement based on struct and fields if any are recieved.
func GenerateNewSelectStatement(table interface{}, fields []string) string {
	tableName := utils.GetTableName(table)
	if len(fields) > 0 {
		return fmt.Sprintf("SELECT (%s) FROM %s;", strings.Join(fields, ","), tableName)

	}
	return fmt.Sprintf("SELECT * FROM %s;", tableName)
}

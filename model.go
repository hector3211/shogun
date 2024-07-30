package shogun

import (
	"fmt"
	"reflect"
	"shogun/utils"
	"strings"
)

type ModelBuilder struct {
	Name    string
	Columns [][]string
}

var emptyModel ModelBuilder

func NewModel(table interface{}) *ModelBuilder {
	t := reflect.TypeOf(table)
	tableName := strings.ToLower(t.Name())

	if t.Kind() != reflect.Struct {
		return &emptyModel
	}

	var columnTags [][]string
	// fmt.Printf("Table name: %s\n", tableName)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		ormTag := field.Tag.Get("orm")
		sqlType := utils.ToSqlTypes(field.Type.Name())

		if ormTag == "id" {
			columnTags = append(columnTags, []string{fmt.Sprintf("%s %s", ormTag, utils.PRIMARYKEY)})
		} else {
			columnTags = append(columnTags, []string{fmt.Sprintf("%s %s", ormTag, sqlType)})
		}
	}

	return &ModelBuilder{
		Name:    tableName,
		Columns: columnTags,
	}
}

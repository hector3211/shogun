package shogun

import (
	"fmt"
	"reflect"
	"shogun/utils"
	"strings"
)

type Model struct {
	ModelType reflect.Type
	Name      string
	OrmTags   []string
}

var emptyModel Model

func NewModel(table interface{}) *Model {
	t := reflect.TypeOf(table)
	tableName := strings.ToLower(t.Name())

	if t.Kind() != reflect.Struct {
		return &emptyModel
	}

	var columnTags []string
	// fmt.Printf("Table name: %s\n", tableName)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		ormTag := field.Tag.Get("orm")
		sqlType := utils.ToSqlTypes(field.Type.Name())

		if ormTag == "id" {
			columnTags = append(columnTags, fmt.Sprintf("%s %s", ormTag, utils.PRIMARYKEY))
		} else {
			columnTags = append(columnTags, fmt.Sprintf("%s %s", ormTag, sqlType))
		}
	}

	return &Model{
		ModelType: t,
		Name:      tableName,
		OrmTags:   columnTags,
	}
}

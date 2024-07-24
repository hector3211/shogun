package utils

import (
	"reflect"
	"strings"
)

type SqlType string

const (
	PRIMARYKEY SqlType = "PRIMARYKEY"
	INTEGER    SqlType = "INTEGER"
	VARCHAR    SqlType = "VARCHAR"
	TEXT       SqlType = "TEXT"
	BOOL       SqlType = "BOOL"
	NULL       SqlType = "NULL"
)

type Driver string

const (
	Postgres Driver = "postgres"
	Sqlite   Driver = "sqlite3"
)

func ToSqlTypes(goType string) SqlType {
	var machedType SqlType
	// tok := reflect.TypeOf(v).String()
	// fmt.Printf("type in string :%s\n", tok)
	switch goType {
	case "int":
		machedType = INTEGER
	case "uint":
		machedType = INTEGER
	case "string":
		machedType = TEXT
	case "bool":
		machedType = BOOL
	default:
		machedType = NULL
	}
	return machedType
}

func MatchDriver(d Driver) string {
	var driver string
	switch d {
	case Postgres:
		driver = "postgres"
	case Sqlite:
		driver = "sqlite3"
	}

	return driver
}

// type Table interface {
// 	TableName() string
// }

func GetTableName(table interface{}) string {
	t := reflect.TypeOf(table)

	if t.Kind() != reflect.Struct {
		return ""
	}

	tableName := strings.ToLower(t.Name())
	return tableName
}

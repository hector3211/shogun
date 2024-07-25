package shogun

import (
	"fmt"
	"strings"
)

type CreateTableBuilder struct {
	Driver
	Verb        string
	ifNotExists bool
	Name        string
	Columns     [][]string
}

func NewCreateTableBuilder() *CreateTableBuilder {
	return &CreateTableBuilder{
		Driver:      DefaultDriver,
		Verb:        "CREATE TABLE",
		ifNotExists: false,
	}
}

func (ct *CreateTableBuilder) CreaetTable(tableName string) *CreateTableBuilder {
	ct.Name = tableName
	return ct
}

func (ct *CreateTableBuilder) IfNotExists() *CreateTableBuilder {
	ct.ifNotExists = true
	return ct
}

func (ct *CreateTableBuilder) Define(val ...string) *CreateTableBuilder {
	ct.Columns = append(ct.Columns, val)
	return ct
}

func (ct *CreateTableBuilder) String() string {
	var stmt string
	if ct.ifNotExists {
		stmt += fmt.Sprintf("%s IF NOT EXISTS %s ", ct.Verb, ct.Name)
	} else {
		stmt += fmt.Sprintf("%s %s", ct.Verb, ct.Name)
	}
	if len(ct.Columns) > 0 {
		stmt += "("
		for i := 0; i < len(ct.Columns); i++ {
			col := ct.Columns[i]
			stmt += strings.Join(col, " ")
			if i != len(ct.Columns)-1 {
				stmt += ","
			}
		}
		stmt += ")"
	}
	stmt += ";"
	return stmt
}

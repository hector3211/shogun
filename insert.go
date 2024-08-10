package shogun

import (
	"fmt"
	"strings"
)

type InsertBuilder struct {
	driver    Driver
	action    string
	tableName string
	columns   []string
	values    []interface{}
}

// Creates a new instance of the InsertBuilder struct
func NewInsertBuilder() *InsertBuilder {
	return DefaultDriver.NewInsertBuilder()
}

func newInsertBuilder() *InsertBuilder {
	return &InsertBuilder{
		action:  "INSERT",
		columns: make([]string, 0),
	}
}

func Insert(tableName string) *InsertBuilder {
	return NewInsertBuilder().Table(tableName)
}

// Sets table name
func (i *InsertBuilder) Table(tableName string) *InsertBuilder {
	i.tableName = tableName
	return i
}

func (i *InsertBuilder) Cols(columns ...string) *InsertBuilder {
	i.columns = columns
	return i
}

func (i *InsertBuilder) Vals(values ...interface{}) *InsertBuilder {
	i.values = values
	return i
}

func (i InsertBuilder) String() string {
	return i.Build()
}

func (i *InsertBuilder) Build() string {
	buf := newStringBuilder()
	buf.WriteLeadingString(fmt.Sprintf("%s INTO %s ", i.action, i.tableName))

	buf.WriteString("(")
	if len(i.columns) > 1 {
		buf.WriteString(strings.Join(i.columns, ","))
	} else {
		buf.WriteString(i.columns[0])
	}
	buf.WriteString(")")

	buf.WriteLeadingString("VALUES ")
	if len(i.values) > 0 {
		buf.WriteString("(")
		for j := 0; j < len(i.values); j++ {
			val := i.values[j]
			switch v := val.(type) {
			case string:
				buf.WriteString(fmt.Sprintf("'%s'", v))
			case int:
				buf.WriteString(fmt.Sprintf("%d", v))
			case float32:
				buf.WriteString(fmt.Sprintf("%f", v))
			case bool:
				buf.WriteString(strings.ToUpper(fmt.Sprintf("%v", v)))
			}

			if j < len(i.values)-1 {
				buf.WriteString(",")
			}
		}
		buf.WriteString(")")
	}
	buf.WriteString(";")

	return buf.String()
}

func (i *InsertBuilder) SetDriver(sqlDriver Driver) *InsertBuilder {
	i.driver = sqlDriver
	return i
}

func (i InsertBuilder) GetDriver() Driver {
	return i.driver
}

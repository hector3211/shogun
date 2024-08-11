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

// Creates a new instance of the InsertBuilder struct
func newInsertBuilder() *InsertBuilder {
	return &InsertBuilder{
		action:  "INSERT",
		columns: make([]string, 0),
	}
}

func Insert(tableName string) *InsertBuilder {
	return NewInsertBuilder().Insert(tableName)
}

// Sets table name
func (i *InsertBuilder) Insert(tableName string) *InsertBuilder {
	i.tableName = tableName
	return i
}

// Loads up fields that will be targeted
func (i *InsertBuilder) Cols(columns ...string) *InsertBuilder {
	i.columns = columns
	return i
}

// Sets the values
func (i *InsertBuilder) Vals(values ...interface{}) *InsertBuilder {
	i.values = values
	return i
}

// Returns back the query in string format
func (i InsertBuilder) String() string {
	return i.Build()
}

// Builds out the final query
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

// Sets a new driver
func (i *InsertBuilder) SetDriver(sqlDriver Driver) *InsertBuilder {
	i.driver = sqlDriver
	return i
}

// Returns current driver being used
func (i InsertBuilder) GetDriver() Driver {
	return i.driver
}

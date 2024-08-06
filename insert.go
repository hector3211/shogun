package shogun

import (
	"fmt"
	"strings"
)

type InsertBuilder struct {
	Action    string
	TableName string
	Columns   []string
	Values    []interface{}
}

// Creates a new instance of the InsertBuilder struct
func NewInsertBuilder() *InsertBuilder {
	return &InsertBuilder{
		Action:  "INSERT",
		Columns: make([]string, 0),
	}
}

// Sets table name
func (i *InsertBuilder) Table(tableName string) *InsertBuilder {
	i.TableName = tableName
	return i
}

func (i *InsertBuilder) Cols(columns ...string) *InsertBuilder {
	i.Columns = columns
	return i
}

func (i *InsertBuilder) Vals(values ...interface{}) *InsertBuilder {
	i.Values = values
	return i
}

func (i *InsertBuilder) Build() string {
	buf := newStringBuilder()
	buf.WriteLeadingString(fmt.Sprintf("%s INTO %s ", i.Action, i.TableName))

	buf.WriteString("(")
	if len(i.Columns) > 1 {
		// Yes LPAREN or RPAREN
		buf.WriteString(strings.Join(i.Columns, ","))
	} else {
		buf.WriteString(i.Columns[0])
	}
	buf.WriteString(")")

	buf.WriteLeadingString("VALUES ")
	if len(i.Values) > 0 {
		buf.WriteString("(")
		for j := 0; j < len(i.Values); j++ {
			val := i.Values[j]
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

			if j < len(i.Values)-1 {
				buf.WriteString(",")
			}
		}
		buf.WriteString(")")
	}
	buf.WriteString(";")

	return buf.String()
}

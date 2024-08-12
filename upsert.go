package shogun

import (
	"fmt"
	"strings"
)

type UpsertBuilder struct {
	driver        Driver
	action        string
	tableName     string
	columns       []string
	values        []interface{}
	conflictField string
	updateField   string
	updatedValue  interface{}
}

func NewUpsertBuilder() *UpsertBuilder {
	return DefaultDriver.NewUpsertBuilder()
}

func newUpsertBuilder() *UpsertBuilder {
	return &UpsertBuilder{
		driver: DefaultDriver,
		action: "INSERT",
	}
}

func Upsert(tableName string) *UpsertBuilder {
	return newUpsertBuilder().Upsert(tableName)
}

func (u *UpsertBuilder) Upsert(tableName string) *UpsertBuilder {
	u.tableName = tableName
	return u
}

// Loads up fields that will be targeted
func (u *UpsertBuilder) Cols(columns ...string) *UpsertBuilder {
	u.columns = columns
	return u
}

// Loads up fields that will be targeted
func (u *UpsertBuilder) Vals(values ...interface{}) *UpsertBuilder {
	u.values = values
	return u
}

func (u *UpsertBuilder) ConflictOn(field string) *UpsertBuilder {
	u.conflictField = field
	return u
}

func (u *UpsertBuilder) Update(field string, updatedValue interface{}) *UpsertBuilder {
	u.updateField = field
	u.updatedValue = updatedValue
	return u
}

func (u *UpsertBuilder) String() string {
	return u.Build()
}

func (u *UpsertBuilder) Build() string {
	buf := newStringBuilder()
	buf.WriteLeadingString(fmt.Sprintf("%s INTO %s ", u.action, u.tableName))
	buf.WriteString("(")
	if len(u.columns) > 1 {
		buf.WriteString(strings.Join(u.columns, ","))
	} else {
		buf.WriteString(u.columns[0])
	}
	buf.WriteString(")")

	buf.WriteLeadingString("VALUES ")
	if len(u.values) > 0 {
		buf.WriteString("(")
		for j := 0; j < len(u.values); j++ {
			val := u.values[j]
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

			if j < len(u.values)-1 {
				buf.WriteString(",")
			}
		}
		buf.WriteString(")")
	}

	buf.WriteLeadingString(fmt.Sprintf("ON CONFLICT(%s) ", u.conflictField))
	buf.WriteLeadingString(fmt.Sprintf("DO UPDATE SET %s = ", u.updateField))
	switch v := u.updatedValue.(type) {
	case string:
		buf.WriteString(fmt.Sprintf("'%s'", v))
	case int:
		buf.WriteString(fmt.Sprintf("%d", v))
	case float32:
		buf.WriteString(fmt.Sprintf("%f", v))
	case bool:
		buf.WriteString(strings.ToUpper(fmt.Sprintf("%v", v)))
	}

	buf.WriteString(";")
	return buf.String()
}

// Sets a new driver
func (i *UpsertBuilder) SetDriver(sqlDriver Driver) *UpsertBuilder {
	i.driver = sqlDriver
	return i
}

// Returns current driver being used
func (i UpsertBuilder) GetDriver() Driver {
	return i.driver
}

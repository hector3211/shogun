package shogun

import (
	"fmt"
	"strings"
)

type DeleteBuilder struct {
	driver     Driver
	tableName  string
	conditions Conditions
}

// creates new instance of DeleteBuilder
func NewDeleteBuilder() *DeleteBuilder {
	return DefaultDriver.NewDeleteBuilder()
}

// creates new instance of DeleteBuilder
func newDeleteBuilder() *DeleteBuilder {
	return &DeleteBuilder{
		driver:     "",
		tableName:  "",
		conditions: [][]string{},
	}
}

// Sets up table that will be targeted
func Delete(tableName string) *DeleteBuilder {
	return newDeleteBuilder().Delete(tableName)
}

// Sets up table that will be targeted
func (d *DeleteBuilder) Delete(tableName string) *DeleteBuilder {
	d.tableName = tableName
	return d
}

// Sets the fields that the query will filter from
func (d *DeleteBuilder) Where(conditions ...string) *DeleteBuilder {
	d.conditions = append(d.conditions, conditions)
	return d
}

// Returns back the query in string format
func (d *DeleteBuilder) String() string {
	return d.Build()
}

// Builds out the final query
func (d *DeleteBuilder) Build() string {
	buf := newStringBuilder()
	buf.WriteLeadingString("DELETE ")

	if d.tableName != "" {
		buf.WriteString(fmt.Sprintf("%s", d.tableName))
	}

	if len(d.conditions) > 0 {
		buf.WriteLeadingString("WHERE ")
		for _, args := range d.conditions {
			buf.WriteString(fmt.Sprintf("%s", strings.Join(args, " ")))
		}
	}

	buf.WriteString(";")

	return buf.String()
}

// Sets a new driver
func (d *DeleteBuilder) SetDriver(sqlDriver Driver) *DeleteBuilder {
	d.driver = sqlDriver
	return d
}

// Returns current driver being used
func (d DeleteBuilder) GetDriver() Driver {
	return d.driver
}

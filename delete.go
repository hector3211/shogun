package shogun

import (
	"fmt"
	"strings"
)

type DeleteQuery interface {
	Delete(tableName string) *DeleteBuilder
	Where(conditions ...string) *DeleteBuilder
	String() string
	Build() string
	SetDriver(sqlDriver Driver) *DeleteBuilder
	GetDriver() string
}

type DeleteBuilder struct {
	driver    Driver
	tableName string
	Conditions
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
		Conditions: make(Conditions, 0),
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
	d.Conditions = append(d.Conditions, conditions)
	return d
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

	if len(d.Conditions) > 0 {
		buf.WriteLeadingString("WHERE ")
		for _, args := range d.Conditions {
			buf.WriteString(fmt.Sprintf("%s", strings.Join(args, " ")))
		}
	}

	buf.WriteString(";")

	return buf.String()
}

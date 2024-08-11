package shogun

import (
	"fmt"
	"strings"
)

type CreateTableBuilder struct {
	driver      Driver
	action      string
	ifNotExists bool
	name        string
	columns     [][]string
}

// creates new instance of CreateTableBuilder
func NewCreateTableBuilder() *CreateTableBuilder {
	return DefaultDriver.NewCreateBuilder()
}

// creates new instance of CreateTableBuilder
func newCreateTableBuilder() *CreateTableBuilder {
	return &CreateTableBuilder{
		driver:      DefaultDriver,
		action:      "CREATE TABLE",
		ifNotExists: false,
	}
}

// Sets the table name for a new table
func CreatTable(tableName string) *CreateTableBuilder {
	return NewCreateTableBuilder().CreateTable(tableName)
}

// Sets the table name for a new table
func (c *CreateTableBuilder) CreateTable(tableName string) *CreateTableBuilder {
	c.name = tableName
	return c
}

// Adds 'IF NOT EXISTS' statement to the query
func (c *CreateTableBuilder) IfNotExists() *CreateTableBuilder {
	c.ifNotExists = true
	return c
}

// Sets table fields with their respected types
func (c *CreateTableBuilder) Define(val ...string) *CreateTableBuilder {
	c.columns = append(c.columns, val)
	return c
}

// Returns query in a string format
func (c *CreateTableBuilder) String() string {
	return c.Build()
}

// Builds out the final query
func (c *CreateTableBuilder) Build() string {
	buf := newStringBuilder()
	if c.ifNotExists {
		buf.WriteLeadingString(fmt.Sprintf("%s ", c.action))
		buf.WriteString(fmt.Sprintf("IF NOT EXISTS %s ", c.name))
	} else {
		buf.WriteString(fmt.Sprintf("%s %s ", c.action, c.name))
	}
	if len(c.columns) > 0 {
		buf.WriteString("(")
		for i := 0; i < len(c.columns); i++ {
			col := c.columns[i]
			buf.WriteString(strings.Join(col, " "))
			if i != len(c.columns)-1 {
				buf.WriteString(",")
			}
		}
		buf.WriteString(")")
	}
	buf.WriteString(";")
	return buf.builder.String()
}

func (c *CreateTableBuilder) SetDriver(sqlDriver Driver) *CreateTableBuilder {
	c.driver = sqlDriver
	return c
}

func (c CreateTableBuilder) GetDriver() Driver {
	return c.driver
}

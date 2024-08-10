package shogun

import (
	"fmt"
	"strings"
)

type CreateTableBuilder struct {
	Driver
	Action      string
	ifNotExists bool
	Name        string
	Columns     [][]string
}

func NewCreateTableBuilder() *CreateTableBuilder {
	return &CreateTableBuilder{
		Driver:      DefaultDriver,
		Action:      "CREATE TABLE",
		ifNotExists: false,
	}
}

func CreatTable(tableName string) *CreateTableBuilder {
	return NewCreateTableBuilder().CreateTable(tableName)
}

func (c *CreateTableBuilder) CreateTable(tableName string) *CreateTableBuilder {
	c.Name = tableName
	return c
}

func (c *CreateTableBuilder) IfNotExists() *CreateTableBuilder {
	c.ifNotExists = true
	return c
}

func (c *CreateTableBuilder) Define(val ...string) *CreateTableBuilder {
	c.Columns = append(c.Columns, val)
	return c
}

func (c *CreateTableBuilder) String() string {
	buf := newStringBuilder()
	if c.ifNotExists {
		buf.WriteLeadingString(fmt.Sprintf("%s ", c.Action))
		buf.WriteString(fmt.Sprintf("IF NOT EXISTS %s ", c.Name))
	} else {
		buf.WriteString(fmt.Sprintf("%s %s ", c.Action, c.Name))
	}
	if len(c.Columns) > 0 {
		buf.WriteString("(")
		for i := 0; i < len(c.Columns); i++ {
			col := c.Columns[i]
			buf.WriteString(strings.Join(col, " "))
			if i != len(c.Columns)-1 {
				buf.WriteString(",")
			}
		}
		buf.WriteString(")")
	}
	buf.WriteString(";")
	return buf.builder.String()
}

func (c *CreateTableBuilder) Build() string {
	buf := newStringBuilder()
	if c.ifNotExists {
		buf.WriteLeadingString(fmt.Sprintf("%s ", c.Action))
		buf.WriteString(fmt.Sprintf("IF NOT EXISTS %s ", c.Name))
	} else {
		buf.WriteString(fmt.Sprintf("%s %s ", c.Action, c.Name))
	}
	if len(c.Columns) > 0 {
		buf.WriteString("(")
		for i := 0; i < len(c.Columns); i++ {
			col := c.Columns[i]
			buf.WriteString(strings.Join(col, " "))
			if i != len(c.Columns)-1 {
				buf.WriteString(",")
			}
		}
		buf.WriteString(")")
	}
	buf.WriteString(";")
	return buf.builder.String()
}

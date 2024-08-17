package shogun

import (
	"fmt"
	"strings"
)

type UpdateQuery interface {
	Update(tableName string) *UpdateBuilder
	Set(values ...string) *UpdateBuilder
	Where(conditions ...string) *UpdateBuilder
	String() string
	Build() string
	SetDriver(sqlDriver Driver) *UpdateBuilder
	GetDriver() string
}

type UpdateBuilder struct {
	driver  Driver
	action  string
	table   string
	setCond Conditions
	Conditions
}

func NewUpdateBuilder() *UpdateBuilder {
	return DefaultDriver.NewUpdateBuilder()
}

func newUpdateBuilder() *UpdateBuilder {
	return &UpdateBuilder{
		action: "UPDATE",
	}
}

// Sets the name of the table being updated
func Update(tableName string) *UpdateBuilder {
	return NewUpdateBuilder().Update(tableName)
}

// Sets the name of the table being updated
func (u *UpdateBuilder) Update(tableName string) *UpdateBuilder {
	u.table = tableName
	return u
}

// Loads up the new values that will be added
func (u *UpdateBuilder) Set(values ...string) *UpdateBuilder {
	u.setCond = append(u.setCond, values)
	return u
}

// Sets the fields that will be updating
func (u *UpdateBuilder) Where(conditions ...string) *UpdateBuilder {
	u.Conditions = append(u.Conditions, conditions)
	return u
}

// Builds out the final query
func (u *UpdateBuilder) Build() string {
	buf := newStringBuilder()
	buf.WriteLeadingString(fmt.Sprintf("%s %s", u.action, u.table))

	if len(u.setCond) > 0 {
		buf.WriteLeadingString("SET ")
		for _, args := range u.setCond {
			buf.WriteString(fmt.Sprintf("%s", strings.Join(args, " ")))
		}
	}

	if len(u.Conditions) > 0 {
		buf.WriteLeadingString("WHERE ")
		for _, args := range u.Conditions {
			buf.WriteString(fmt.Sprintf("%s", strings.Join(args, " ")))
		}
	}
	buf.WriteString(";")
	return buf.String()
}

// Sets a new driver
func (s *UpdateBuilder) SetDriver(sqlDriver Driver) *UpdateBuilder {
	s.driver = sqlDriver
	return s
}

// Returns current driver being used
func (s UpdateBuilder) GetDriver() Driver {
	return s.driver
}

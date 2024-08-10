package shogun

import (
	"fmt"
	"strings"
)

type UpdateBuilder struct {
	action  string
	table   string
	setCond Conditions
	Conditions
}

func NewUpdateBuilder() *UpdateBuilder {
	return &UpdateBuilder{
		action: "UPDATE",
	}
}

func Update(tableName string) *UpdateBuilder {
	return NewUpdateBuilder().Update(tableName)
}

func (u *UpdateBuilder) Update(tableName string) *UpdateBuilder {
	u.table = tableName
	return u
}

func (u *UpdateBuilder) Set(values ...string) *UpdateBuilder {
	u.setCond = append(u.setCond, values)
	return u
}

func (u *UpdateBuilder) Where(conditions ...string) *UpdateBuilder {
	u.Conditions = append(u.Conditions, conditions)
	return u
}

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

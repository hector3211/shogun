package shogun

import (
	"fmt"
	"strings"
)

type UpdateBuilder struct {
	Action  string
	Table   string
	SetCond Conditions
	Conditions
}

func NewUpdateBuilder() *UpdateBuilder {
	return &UpdateBuilder{
		Action: "UPDATE",
	}
}

func (u *UpdateBuilder) Update(tableName string) *UpdateBuilder {
	u.Table = tableName
	return u
}

func (u *UpdateBuilder) Set(values ...string) *UpdateBuilder {
	u.SetCond = append(u.SetCond, values)
	return u
}

func (u *UpdateBuilder) Where(conditions ...string) *UpdateBuilder {
	u.Conditions = append(u.Conditions, conditions)
	return u
}

func (u *UpdateBuilder) Build() string {
	buf := newStringBuilder()
	buf.WriteLeadingString(fmt.Sprintf("%s %s", u.Action, u.Table))

	if len(u.SetCond) > 0 {
		buf.WriteLeadingString("SET")
		for _, args := range u.SetCond {
			buf.WriteString(fmt.Sprintf(" %s", strings.Join(args, " ")))
		}
	}

	if len(u.Conditions) > 0 {
		buf.WriteLeadingString("WHERE")
		for _, args := range u.Conditions {
			buf.WriteString(fmt.Sprintf(" %s", strings.Join(args, " ")))
		}
	}
	buf.WriteString(";")
	return buf.String()
}

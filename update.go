package shogun

import (
	"fmt"
	"strings"
)

type UpdateBuilder struct {
	Action    string
	Table     string
	SetCond   [][]interface{}
	WhereCond [][]interface{}
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

func (u *UpdateBuilder) Set(values ...interface{}) *UpdateBuilder {
	u.SetCond = append(u.SetCond, values)
	return u
}

func (u *UpdateBuilder) Where(conditions ...interface{}) *UpdateBuilder {
	u.WhereCond = append(u.WhereCond, conditions)
	return u
}

func (u *UpdateBuilder) Build() string {
	buf := newStringBuilder()
	buf.WriteLeadingString(fmt.Sprintf("%s ", u.Action))

	//NOTE:  Not sure about this double for loop
	for _, setCond := range u.SetCond {
		var set string
		for _, token := range setCond {
			switch token.(type) {
			case string:
				set += fmt.Sprintf("%s", token)
			case int:
				set += fmt.Sprintf("%d", token)
			case float32:
				set += fmt.Sprintf("%f", token)
			case bool:
				set += strings.ToUpper(fmt.Sprintf("%v", token))
			}
		}
	}

	return buf.String()
}

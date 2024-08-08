package shogun

import (
	"fmt"
	"strings"
)

type Conditions [][]string

// func (w *WhereCond) Where(conditions ...string) *WhereCond {
// 	w.Cond = append(w.Cond, conditions)
// 	return w
// }

func Equal(field string, value interface{}) string {
	return stringifyStatement(field, "=", value)
}

func NotEqual(field string, value interface{}) string {
	return stringifyStatement(field, "!=", value)
}

func LessThan(field string, value interface{}) string {
	return stringifyStatement(field, "<", value)
}

func GreaterThan(field string, value interface{}) string {
	return stringifyStatement(field, ">", value)
}

func And() string {
	buf := newStringBuilder()
	buf.WriteString("AND")
	return buf.String()
}

func Or() string {
	buf := newStringBuilder()
	buf.WriteString("OR")
	return buf.String()
}

func stringifyStatement(field, action string, value interface{}) string {
	var statement string
	switch value.(type) {
	case int:
		statement = fmt.Sprintf("%s %s %d", field, action, value)
	case float32:
		statement = fmt.Sprintf("%s %s %f", field, action, value)
	case string:
		statement = fmt.Sprintf("%s %s '%s'", field, action, value)
	case bool:
		strBool := fmt.Sprintf("%v", value)
		statement = fmt.Sprintf("%s %s '%s'", field, action, strings.ToUpper(strBool))
	default:
		statement = fmt.Sprintf("%s %s %v", field, action, value)
	}

	return statement
}

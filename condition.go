package shogun

import (
	"fmt"
	"strings"
)

type Conditions [][]string
type ConditionToken int

const (
	EQUAL ConditionToken = iota
	NOTEQUAL
	LESSTHAN
	GREATERTHAN
)

func (c ConditionToken) String() string {
	switch c {
	case EQUAL:
		return "="
	case NOTEQUAL:
		return "!="
	case LESSTHAN:
		return "<"
	case GREATERTHAN:
		return ">"
	default:
		return ""
	}
}

func Equal(field string, value interface{}) string {
	return stringifyStatement(field, EQUAL, value)
}

func NotEqual(field string, value interface{}) string {
	return stringifyStatement(field, NOTEQUAL, value)
}

func LessThan(field string, value interface{}) string {
	return stringifyStatement(field, LESSTHAN, value)
}

func GreaterThan(field string, value interface{}) string {
	return stringifyStatement(field, GREATERTHAN, value)
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

func stringifyStatement(field string, condition ConditionToken, value interface{}) string {
	var statement string
	switch value.(type) {
	case int:
		statement = fmt.Sprintf("%s %s %d", field, condition.String(), value)
	case float32:
		statement = fmt.Sprintf("%s %s %f", field, condition.String(), value)
	case string:
		statement = fmt.Sprintf("%s %s '%s'", field, condition.String(), value)
	case bool:
		strBool := fmt.Sprintf("%v", value)
		statement = fmt.Sprintf("%s %s %s", field, condition.String(), strings.ToUpper(strBool))
	default:
		statement = fmt.Sprintf("%s %s %v", field, condition.String(), value)
	}

	return statement
}

package shogun

import (
	"fmt"
	"strings"
)

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

type Conditions [][]string

func (c Conditions) Equal(field string, value interface{}) string {
	return stringifyStatement(field, EQUAL, value)
}

func (c Conditions) NotEqual(field string, value interface{}) string {
	return stringifyStatement(field, NOTEQUAL, value)
}

func (c Conditions) LessThan(field string, value interface{}) string {
	return stringifyStatement(field, LESSTHAN, value)
}

func (c Conditions) GreaterThan(field string, value interface{}) string {
	return stringifyStatement(field, GREATERTHAN, value)
}

func (c Conditions) And() string {
	buf := newStringBuilder()
	buf.WriteString("AND")
	return buf.String()
}

func (c Conditions) Or() string {

	buf := newStringBuilder()
	buf.WriteString("OR")
	return buf.String()
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
	buf := newStringBuilder()
	switch value.(type) {
	case int, float32:
		buf.WriteString(fmt.Sprintf("%s %s %d", field, condition.String(), value))
	case string:
		buf.WriteString(fmt.Sprintf("%s %s '%s'", field, condition.String(), value))
	case bool:
		strBool := fmt.Sprintf("%v", value)
		buf.WriteString(fmt.Sprintf("%s %s %s", field, condition.String(), strings.ToUpper(strBool)))
	default:
		buf.WriteString(fmt.Sprintf("%s %s %v", field, condition.String(), value))
	}

	return buf.String()
}

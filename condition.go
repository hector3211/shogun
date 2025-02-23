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
	BETWEEN
	ISNULL
	NOTNULL
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
	case BETWEEN:
		return "BETWEEN"
	case ISNULL:
		return "IS NULL"
	case NOTNULL:
		return "IS NOT NULL"
	default:
		return ""
	}
}

type CalculationToken int

const (
	SUM CalculationToken = iota
	COUNT
	AVG
	MAX
	MIN
)

func (c CalculationToken) String() string {
	switch c {
	case SUM:
		return "SUM"
	case COUNT:
		return "COUNT"
	case AVG:
		return "AVG"
	case MAX:
		return "MAX"
	case MIN:
		return "MIN"
	default:
		return ""
	}
}

type Conditions [][]string

func (c Conditions) Equal(field string, value interface{}) string {
	return stringifyStatement(field, EQUAL, value, nil)
}

func (c Conditions) NotEqual(field string, value interface{}) string {
	return stringifyStatement(field, NOTEQUAL, value, nil)
}

func (c Conditions) LessThan(field string, value interface{}) string {
	return stringifyStatement(field, LESSTHAN, value, nil)
}

func (c Conditions) GreaterThan(field string, value interface{}) string {
	return stringifyStatement(field, GREATERTHAN, value, nil)
}

func (c Conditions) Between(field string, value interface{}) string {
	return stringifyStatement(field, BETWEEN, value, nil)
}

func (c Conditions) IsNull(field string) string {
	return stringifyStatement(field, ISNULL, nil, nil)
}

func (c Conditions) IsNOTNull(field string) string {
	return stringifyStatement(field, NOTNULL, nil, nil)
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
	return stringifyStatement(field, EQUAL, value, nil)
}

func NotEqual(field string, value interface{}) string {
	return stringifyStatement(field, NOTEQUAL, value, nil)
}

func LessThan(field string, value interface{}) string {
	return stringifyStatement(field, LESSTHAN, value, nil)
}

func GreaterThan(field string, value interface{}) string {
	return stringifyStatement(field, GREATERTHAN, value, nil)
}

func Between(field string, value interface{}) string {
	return stringifyStatement(field, BETWEEN, value, nil)
}

func IsNull(field string) string {
	return stringifyStatement(field, ISNULL, nil, nil)
}

func IsNotNull(field string) string {
	return stringifyStatement(field, NOTNULL, nil, nil)
}

func Sum(field string, condition ConditionToken, value interface{}) string {
	key := SUM
	return stringifyStatement(field, condition, value, &key)
}

func Count(field string, condition ConditionToken, value interface{}) string {
	key := COUNT
	return stringifyStatement(field, condition, value, &key)
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

func stringifyStatement(field string, condition ConditionToken, value interface{}, calculation *CalculationToken) string {
	buf := newStringBuilder()

	if calculation == nil {
		switch value.(type) {
		case int, float32:
			buf.WriteString(fmt.Sprintf("%s %s %d", field, condition.String(), value))
		case string:
			buf.WriteString(fmt.Sprintf("%s %s '%s'", field, condition.String(), value))
		case bool:
			strBool := fmt.Sprintf("%v", value)
			buf.WriteString(fmt.Sprintf("%s %s %s", field, condition.String(), strings.ToUpper(strBool)))
		default:
			buf.WriteString(fmt.Sprintf("%s %s", field, condition.String()))
		}
	} else {
		switch value.(type) {
		case int, float32:
			buf.WriteString(fmt.Sprintf("%s(%s) %s %d", calculation, field, condition.String(), value))
		case string:
			buf.WriteString(fmt.Sprintf("%s(%s) %s '%s'", calculation, field, condition.String(), value))
		case bool:
			strBool := fmt.Sprintf("%v", value)
			buf.WriteString(fmt.Sprintf("%s(%s) %s %s", calculation, field, condition.String(), strings.ToUpper(strBool)))
		default:
			buf.WriteString(fmt.Sprintf("%s(%s) %s %v", calculation, field, condition.String(), value))
		}
	}

	return buf.String()
}

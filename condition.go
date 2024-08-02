package shogun

import (
	"fmt"
	"reflect"
)

type Condition struct {
	Args []string
}

func NewCondition() *Condition {
	return &Condition{
		Args: []string{},
	}
}

func (c *Condition) Equal(field string, value interface{}) string {
	t := reflect.TypeOf(value)
	var stringValue string

	theType := t.Kind()
	switch theType {
	case reflect.Int:
		stringValue = fmt.Sprintf(" %v ", value)
	case reflect.String:
		stringValue = fmt.Sprintf(" %s ", value)
	default:
		stringValue = fmt.Sprintf(" %s ", value)

	}

	buf := newStringBuilder()
	buf.WriteString(field)
	buf.WriteString(" = ")
	buf.WriteString(stringValue)
	return buf.builder.String()
}

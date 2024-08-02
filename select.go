package shogun

import (
	"fmt"
	"reflect"
	"strings"
)

type SelectBuilder struct {
	Driver
	Tables     [][]string
	Fields     [][]string
	WhereConds [][]string
	And        bool
	Or         bool
}

func NewCreateSelectBuilder() *SelectBuilder {
	return &SelectBuilder{
		Driver:     DefaultDriver,
		Tables:     make([][]string, 0),
		Fields:     make([][]string, 0),
		WhereConds: make([][]string, 0),
	}
}

func (s *SelectBuilder) Select(columns ...string) *SelectBuilder {
	s.Fields = append(s.Fields, columns)
	return s
}

func (s *SelectBuilder) From(table ...string) *SelectBuilder {
	s.Tables = append(s.Tables, table)
	return s
}

func (s *SelectBuilder) Where(conditions ...string) *SelectBuilder {
	s.WhereConds = append(s.WhereConds, conditions)
	return s
}

func Equal(field string, value interface{}) string {
	t := reflect.TypeOf(value)
	var eqStatement string

	theType := t.Kind()
	switch theType {
	case reflect.Int:
		eqStatement = fmt.Sprintf("%s %s %v", field, "=", value)
	case reflect.String:
		eqStatement = fmt.Sprintf("%s %s '%s'", field, "=", value)
	default:
		// could be for bools, could be wrong
		eqStatement = fmt.Sprintf("%s %s %v", field, "=", value)
	}

	return eqStatement
}

func NotEqual(field string, value interface{}) string {
	t := reflect.TypeOf(value)
	var eqStatement string

	theType := t.Kind()
	switch theType {
	case reflect.Int:
		eqStatement = fmt.Sprintf("%s %s %v", field, "!=", value)
	case reflect.String:
		eqStatement = fmt.Sprintf("%s %s '%s'", field, "!=", value)
	default:
		// could be for bools, could be wrong
		eqStatement = fmt.Sprintf("%s %s %v", field, "!=", value)
	}

	return eqStatement
}

func LessThan(field string, value interface{}) string {
	t := reflect.TypeOf(value)
	var lessThanStatement string

	theType := t.Kind()
	switch theType {
	case reflect.Int:
		lessThanStatement = fmt.Sprintf("%s %s %v", field, "<", value)
	case reflect.String:
		lessThanStatement = fmt.Sprintf("%s %s '%s'", field, "<", value)
	default:
		// could be for bools, could be wrong
		lessThanStatement = fmt.Sprintf("%s %s %v", field, "<", value)
	}

	return lessThanStatement
}

func GreaterThan(field string, value interface{}) string {
	t := reflect.TypeOf(value)
	var lessThanStatement string

	theType := t.Kind()
	switch theType {
	case reflect.Int:
		lessThanStatement = fmt.Sprintf("%s %s %v", field, ">", value)
	case reflect.String:
		lessThanStatement = fmt.Sprintf("%s %s '%s'", field, ">", value)
	default:
		// could be for bools, could be wrong
		lessThanStatement = fmt.Sprintf("%s %s %v", field, ">", value)
	}

	return lessThanStatement
}

func And() string {
	return "AND"
}

func Or() string {
	return "OR"
}

func (s *SelectBuilder) Build() string {
	buf := newStringBuilder()
	buf.WriteLeadingString("SELECT ")

	if len(s.Fields[0]) > 1 {
		buf.WriteString("(")
		for _, field := range s.Fields {
			buf.WriteString(strings.Join(field, ","))
		}
		buf.WriteString(")")
	} else {
		if strings.Join(s.Fields[0], "") != "*" {
			buf.WriteString("(")
			for _, field := range s.Fields {
				buf.WriteString(strings.Join(field, ","))
			}
			buf.WriteString(")")
		} else {
			buf.WriteString("*")
		}
	}

	buf.WriteLeadingString("FROM ")

	if len(s.Tables[0]) > 1 {
		buf.WriteString("(")
		for _, table := range s.Tables {
			buf.WriteString(strings.Join(table, ","))
		}
		buf.WriteString(")")
	} else {
		buf.WriteString(strings.Join(s.Tables[0], ""))
	}

	if len(s.WhereConds) > 0 {
		buf.WriteLeadingString("WHERE ")
		for _, args := range s.WhereConds {
			buf.WriteString(strings.Join(args, " "))
		}
	}

	buf.WriteString(";")

	return buf.builder.String()
}

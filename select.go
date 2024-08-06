package shogun

import (
	"fmt"
	"strings"
)

type SelectBuilder struct {
	Driver
	Tables     []string
	Fields     []string
	WhereConds [][]string
}

func NewSelectBuilder() *SelectBuilder {
	return &SelectBuilder{
		Driver:     DefaultDriver,
		Tables:     make([]string, 0),
		Fields:     make([]string, 0),
		WhereConds: make([][]string, 0),
	}
}

func (s *SelectBuilder) Select(columns ...string) *SelectBuilder {
	s.Fields = columns
	return s
}

func (s *SelectBuilder) From(tables ...string) *SelectBuilder {
	s.Tables = tables
	return s
}

func (s *SelectBuilder) Where(conditions ...string) *SelectBuilder {
	s.WhereConds = append(s.WhereConds, conditions)
	return s
}

func Equal(field string, value interface{}) string {
	var eqStatement string

	switch value.(type) {
	case int:
		eqStatement = fmt.Sprintf("%s %s %v", field, "=", value)
	case string:
		eqStatement = fmt.Sprintf("%s %s '%s'", field, "=", value)
	default:
		eqStatement = fmt.Sprintf("%s %s %v", field, "=", value)
	}

	return eqStatement
}

func NotEqual(field string, value interface{}) string {
	var eqStatement string

	switch value.(type) {
	case int:
		eqStatement = fmt.Sprintf("%s %s %v", field, "!=", value)
	case string:
		eqStatement = fmt.Sprintf("%s %s '%s'", field, "!=", value)
	default:
		eqStatement = fmt.Sprintf("%s %s %v", field, "!=", value)
	}

	return eqStatement
}

func LessThan(field string, value interface{}) string {
	var lessThanStatement string

	switch value.(type) {
	case int:
		lessThanStatement = fmt.Sprintf("%s %s %v", field, "<", value)
	case string:
		lessThanStatement = fmt.Sprintf("%s %s '%s'", field, "<", value)
	default:
		lessThanStatement = fmt.Sprintf("%s %s %v", field, "<", value)
	}
	return lessThanStatement
}

func GreaterThan(field string, value interface{}) string {
	var lessThanStatement string

	switch value.(type) {
	case int:
		lessThanStatement = fmt.Sprintf("%s %s %d", field, ">", value)
	case string:
		lessThanStatement = fmt.Sprintf("%s %s '%s'", field, ">", value)
	default:
		lessThanStatement = fmt.Sprintf("%s %s %v", field, ">", value)
	}

	return lessThanStatement
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

func (s *SelectBuilder) Build() string {
	buf := newStringBuilder()
	buf.WriteLeadingString("SELECT ")

	if len(s.Fields) > 1 {
		buf.WriteString(fmt.Sprintf("(%s)", strings.Join(s.Fields, ",")))
	} else {
		if s.Fields[0] != "*" {
			buf.WriteString(fmt.Sprintf("(%s)", s.Fields[0]))
		} else {
			buf.WriteString("*")
		}
	}

	buf.WriteLeadingString("FROM ")

	if len(s.Tables) > 1 {
		buf.WriteString("(")
		buf.WriteString(strings.Join(s.Tables, ","))
		buf.WriteString(")")
	} else {
		buf.WriteString(s.Tables[0])
	}

	if len(s.WhereConds) > 0 {
		buf.WriteLeadingString("WHERE ")
		for _, args := range s.WhereConds {
			buf.WriteString(strings.Join(args, " "))
		}
	}

	buf.WriteString(";")

	return buf.String()
}

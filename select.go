package shogun

import (
	"fmt"
	"strings"
)

type SelectBuilder struct {
	Driver
	Tables []string
	Fields []string
	Conditions
}

func NewSelectBuilder() *SelectBuilder {
	return &SelectBuilder{
		Driver: DefaultDriver,
		Tables: make([]string, 0),
		Fields: make([]string, 0),
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
	s.Conditions = append(s.Conditions, conditions)
	return s
}

func (s *SelectBuilder) Build() string {
	buf := newStringBuilder()
	buf.WriteLeadingString("SELECT")

	if len(s.Fields) > 1 {
		buf.WriteString(fmt.Sprintf(" (%s)", strings.Join(s.Fields, ",")))
	} else {
		if s.Fields[0] != "*" {
			buf.WriteString(fmt.Sprintf(" (%s)", s.Fields[0]))
		} else {
			buf.WriteString(fmt.Sprintf(" %s", "*"))
		}
	}

	buf.WriteLeadingString("FROM")

	if len(s.Tables) > 1 {
		buf.WriteString(fmt.Sprintf(" (%s)", strings.Join(s.Tables, ",")))
	} else {
		buf.WriteString(fmt.Sprintf(" %s", s.Tables[0]))
	}

	if len(s.Conditions) > 0 {
		buf.WriteLeadingString("WHERE")
		for _, args := range s.Conditions {
			buf.WriteString(fmt.Sprintf(" %s", strings.Join(args, " ")))
		}
	}

	buf.WriteString(";")

	return buf.String()
}

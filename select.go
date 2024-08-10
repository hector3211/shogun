package shogun

import (
	"fmt"
	"strings"
)

type SelectBuilder struct {
	driver      Driver
	tables      []string
	fields      []string
	orderFields []string
	conditions  Conditions
	limit       int
	order       string
}

func NewSelectBuilder() *SelectBuilder {
	return DefaultDriver.NewSelectBuilder()
}

func newSelectbuilder() *SelectBuilder {
	return &SelectBuilder{
		driver: DefaultDriver,
		tables: make([]string, 0),
		fields: make([]string, 0),
		limit:  0,
	}
}

func Select(columns ...string) *SelectBuilder {
	return NewSelectBuilder().Select(columns...)
}

func (s *SelectBuilder) TableNames() []string {
	return s.tables
}

func (s *SelectBuilder) Select(columns ...string) *SelectBuilder {
	s.fields = columns
	return s
}

func (s *SelectBuilder) From(tables ...string) *SelectBuilder {
	s.tables = tables
	return s
}

func (s *SelectBuilder) Where(conditions ...string) *SelectBuilder {
	s.conditions = append(s.conditions, conditions)
	return s
}

func (s *SelectBuilder) OrderBy(columns ...string) *SelectBuilder {
	s.orderFields = columns
	return s
}

func (s *SelectBuilder) Asc() *SelectBuilder {
	s.order = "ASC"
	return s
}

func (s *SelectBuilder) Desc() *SelectBuilder {
	s.order = "DESC"
	return s
}

func (s *SelectBuilder) Limit(number int) *SelectBuilder {
	s.limit = number
	return s
}

func (s *SelectBuilder) String() string {
	return s.Build()
}

func (s *SelectBuilder) Build() string {
	buf := newStringBuilder()
	buf.WriteLeadingString("SELECT ")

	if len(s.fields) > 1 {
		buf.WriteString(fmt.Sprintf("(%s)", strings.Join(s.fields, ",")))
	} else {
		if s.fields[0] != "*" {
			buf.WriteString(fmt.Sprintf("(%s)", s.fields[0]))
		} else {
			buf.WriteString(fmt.Sprintf("%s", "*"))
		}
	}

	buf.WriteLeadingString("FROM ")

	if len(s.tables) > 1 {
		buf.WriteString(fmt.Sprintf("(%s)", strings.Join(s.tables, ",")))
	} else {
		buf.WriteString(fmt.Sprintf("%s", s.tables[0]))
	}

	if len(s.conditions) > 0 {
		buf.WriteLeadingString("WHERE ")
		for _, args := range s.conditions {
			buf.WriteString(fmt.Sprintf("%s", strings.Join(args, " ")))
		}
	}

	if len(s.orderFields) > 0 {
		buf.WriteLeadingString("ORDER BY ")
		buf.WriteString(fmt.Sprintf("%s %s", strings.Join(s.orderFields, " "), s.order))
	}

	if s.limit > 0 {
		buf.WriteLeadingString("LIMIT ")
		buf.WriteString(fmt.Sprintf("%d", s.limit))
	}

	buf.WriteString(";")

	return buf.String()
}

func (s *SelectBuilder) SetDriver(sqlDriver Driver) *SelectBuilder {
	s.driver = sqlDriver
	return s
}

func (s SelectBuilder) GetDriver() Driver {
	return s.driver
}

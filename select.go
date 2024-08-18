package shogun

import (
	"fmt"
	"strings"
)

type SelectQuery interface {
	Select(columns ...string) *SelectBuilder
	From(tables ...string) *SelectBuilder
	Where(conditions ...string) *SelectBuilder
	Distinct() *SelectBuilder
	OrderBy(columns ...string) *SelectBuilder
	Asc() *SelectBuilder
	Desc() *SelectBuilder
	Limit(number int) *SelectBuilder
	String() string
	Build() string
	SetDriver(sqlDriver Driver) *SelectBuilder
	GetDriver() string
}

type SelectBuilder struct {
	driver      Driver
	tables      []string
	fields      []string
	orderFields []string
	Conditions
	distinct bool
	limit    int
	groupBy  []string
	order    string
}

// Creates a new instance of the SelectBuilder struct
func NewSelectBuilder() *SelectBuilder {
	return DefaultDriver.NewSelectBuilder()
}

// Creates a new instance of the SelectBuilder struct
func newSelectbuilder() *SelectBuilder {
	return &SelectBuilder{
		driver:      DefaultDriver,
		tables:      make([]string, 0),
		fields:      make([]string, 0),
		orderFields: make([]string, 0),
		groupBy:     make([]string, 0),
		limit:       0,
	}
}

// Sets the fields that query will select returning new instance of SelectBuilder
func Select(columns ...string) *SelectBuilder {
	return NewSelectBuilder().Select(columns...)
}

// Sets the fields that query will select
func (s *SelectBuilder) Select(columns ...string) *SelectBuilder {
	s.fields = columns
	return s
}

func (s *SelectBuilder) Distinct() *SelectBuilder {
	s.distinct = true

	return s
}

// Returns all the table names
func (s *SelectBuilder) TableNames() []string {
	return s.tables
}

// Sets the tables that query will target
func (s *SelectBuilder) From(tables ...string) *SelectBuilder {
	s.tables = tables
	return s
}

// Sets the fields that will be selecting
func (s *SelectBuilder) Where(conditions ...string) *SelectBuilder {
	s.Conditions = append(s.Conditions, conditions)
	return s
}

// TODO: Finish groupby and having
// Sets a Group by to the query
func (s *SelectBuilder) GroupBy(columns ...string) *SelectBuilder {
	s.groupBy = columns
	return s
}

// Sets an order to the query
func (s *SelectBuilder) OrderBy(columns ...string) *SelectBuilder {
	s.orderFields = columns
	return s
}

// Sets order in ascending
func (s *SelectBuilder) Asc() *SelectBuilder {
	s.order = "ASC"
	return s
}

// Sets order in descending
func (s *SelectBuilder) Desc() *SelectBuilder {
	s.order = "DESC"
	return s
}

// Sets the limit for the query
func (s *SelectBuilder) Limit(number int) *SelectBuilder {
	s.limit = number
	return s
}

// Sets a new driver
func (s *SelectBuilder) SetDriver(sqlDriver Driver) *SelectBuilder {
	s.driver = sqlDriver
	return s
}

// Returns current driver being used
func (s SelectBuilder) GetDriver() Driver {
	return s.driver
}

// Returns the query in a string format
func (s *SelectBuilder) String() string {
	return s.Build()
}

// Builds out the final query
func (s *SelectBuilder) Build() string {
	buf := newStringBuilder()
	buf.WriteLeadingString("SELECT ")
	if s.distinct {
		buf.WriteLeadingString("DISTINCT ")
	}

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

	if len(s.Conditions) > 0 {
		buf.WriteLeadingString("WHERE ")
		for _, args := range s.Conditions {
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

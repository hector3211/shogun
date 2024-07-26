package shogun

import "strings"

type SelectBuilder struct {
	Driver
	Tables [][]string
	Args   [][]string
	Where  *string
}

func NewCreateSelectBuilder() *SelectBuilder {
	return &SelectBuilder{
		Driver: DefaultDriver,
		Tables: make([][]string, 0),
		Args:   make([][]string, 0),
	}
}

func (s *SelectBuilder) Select(columns ...string) *SelectBuilder {
	s.Args = append(s.Args, columns)
	return s
}

func (s *SelectBuilder) From(table ...string) *SelectBuilder {
	s.Tables = append(s.Tables, table)
	return s
}

func (s *SelectBuilder) Build() string {
	buf := newStringBuilder()
	buf.WriteLeadingString("SELECT ")

	if len(s.Args[0]) > 1 {
		buf.WriteString("(")
		for _, field := range s.Args {
			buf.WriteString(strings.Join(field, ","))
		}
		buf.WriteString(")")
	} else {
		if strings.Join(s.Args[0], "") != "*" {
			buf.WriteString("(")
			for _, field := range s.Args {
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

	buf.WriteString(";")

	return buf.builder.String()
}

// func (s *SelectBuilder) WhereExpression(expr string) *SelectBuilder {
// 	if len(expr) == 0 {
// 		return s
// 	}
//
// 	if s.Where == nil {
//
// 	}
// 	return s
// }
//
// func (s *SelectBuilder) Equals(equalsString string) *SelectBuilder {
// 	return s
// }

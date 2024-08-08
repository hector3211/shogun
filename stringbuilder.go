package shogun

import "strings"

type StringBuilder struct {
	builder *strings.Builder
}

func newStringBuilder() *StringBuilder {
	return &StringBuilder{
		builder: &strings.Builder{},
	}
}

func (s *StringBuilder) WriteLeadingString(str string) {
	if s.builder.Len() > 0 && s.builder.String()[len(s.builder.String())-1] != ' ' {
		s.builder.WriteString(" ")
	}
	s.builder.WriteString(str)
}

func (s *StringBuilder) WriteString(str string) {
	s.builder.WriteString(str)
}

func (s *StringBuilder) String() string {
	return s.builder.String()
}

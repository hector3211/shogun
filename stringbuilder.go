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

func (sb *StringBuilder) WriteLeadingString(s string) {
	if sb.builder.Len() > 0 {
		sb.builder.WriteString(" ")
	}

	sb.builder.WriteString(s)
}

func (sb *StringBuilder) WriteString(s string) {
	sb.builder.WriteString(s)
}

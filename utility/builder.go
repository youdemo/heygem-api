package util

import "strings"

type Builder struct {
	builder strings.Builder
}

func NewBuilder() *Builder {
	return &Builder{builder: strings.Builder{}}
}

func (b *Builder) Append(s string) *Builder {
	b.builder.WriteString(s)
	return b
}

func (b *Builder) String() string {
	return b.builder.String()
}

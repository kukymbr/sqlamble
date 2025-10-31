package formatter

import (
	"context"
	"go/format"
)

func NewGoFmt() Formatter {
	return &goFmt{}
}

type goFmt struct{}

func (f *goFmt) Format(_ context.Context, content []byte) ([]byte, error) {
	return format.Source(content)
}

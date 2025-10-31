package formatter

import (
	"context"
)

func NewNoop() Formatter {
	return &noop{}
}

type noop struct{}

func (f *noop) Format(_ context.Context, content []byte) ([]byte, error) {
	return content, nil
}

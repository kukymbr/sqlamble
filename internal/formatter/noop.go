package formatter

import (
	"context"
)

func NewNoopFormatter() Formatter {
	return &noop{}
}

type noop struct{}

func (f *noop) Format(_ context.Context, _ string) error {
	return nil
}

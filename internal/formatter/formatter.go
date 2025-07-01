package formatter

import (
	"context"
	"fmt"
)

const (
	GoFmt = "gofmt"
	Noop  = "noop"

	DefaultFormatter = GoFmt
)

func Factory(name string) (Formatter, error) {
	switch name {
	case GoFmt:
		return NewGoFmtFormatter(), nil
	case Noop, "none":
		return NewNoopFormatter(), nil
	}

	return nil, fmt.Errorf("unknown formatter: %s", name)
}

type Formatter interface {
	Format(ctx context.Context, dirPath string) error
}

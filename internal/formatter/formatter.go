package formatter

import (
	"context"
	"fmt"
	"strings"
)

const (
	GoFmt = "gofmt"
	Noop  = "noop"
)

func Factory(name string) (Formatter, error) {
	name = strings.ToLower(name)

	switch name {
	case GoFmt:
		return NewGoFmt(), nil
	case Noop, "none":
		return NewNoop(), nil
	}

	return nil, fmt.Errorf("unknown formatter: %s", name)
}

type Formatter interface {
	Format(ctx context.Context, content []byte) ([]byte, error)
}

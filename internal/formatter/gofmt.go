package formatter

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
)

func NewGoFmtFormatter() Formatter {
	return &goFmt{
		executable: "go",
	}
}

type goFmt struct {
	executable string
}

func (f *goFmt) Format(ctx context.Context, dirPath string) error {
	var errOut bytes.Buffer

	path, err := filepath.Abs(dirPath)
	if err != nil {
		return fmt.Errorf("failed to resolve absolute path for %s: %w", dirPath, err)
	}

	//nolint:gosec
	cmd := exec.CommandContext(ctx, f.executable, "fmt", path+"/...")
	cmd.Stderr = &errOut

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run gofmt: %w (output: %s)", err, errOut.String())
	}

	return nil
}

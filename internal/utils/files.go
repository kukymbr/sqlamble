package utils

import (
	"fmt"
	"os"
)

const filesMode os.FileMode = 0644

func WriteFile(content fmt.Stringer, target string) error {
	if err := os.WriteFile(target, []byte(content.String()), filesMode); err != nil {
		return fmt.Errorf("failed to write file %s: %w", target, err)
	}

	return nil
}

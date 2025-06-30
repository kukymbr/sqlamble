package utils

import (
	"fmt"
	"os"
)

const dirsMode os.FileMode = 0755

// ValidateIsDir checks if path exists and is a directory.
func ValidateIsDir(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("directory '%s': %w", path, err)
	}

	if !stat.IsDir() {
		return fmt.Errorf("'%s' is not a directory", path)
	}

	return nil
}

// EnsureDir creates dir if not exists.
func EnsureDir(path string) error {
	if err := os.MkdirAll(path, dirsMode); err != nil {
		return fmt.Errorf("dir '%s' does not exist and cannot be created: %w", path, err)
	}

	return nil
}

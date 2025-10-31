package utils

import (
	"fmt"
	"os"
)

const (
	filesMode os.FileMode = 0644
	dirsMode  os.FileMode = 0755
)

func WriteFile(content []byte, target string) error {
	if err := os.WriteFile(target, content, filesMode); err != nil {
		return fmt.Errorf("failed to write file %s: %w", target, err)
	}

	return nil
}

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

package utils

import (
	"fmt"
	"regexp"
)

var pxIdentifier = regexp.MustCompile(`(?i)^[a-z]+[a-z0-9_]*$`)

func ValidatePackageName(name string) error {
	if err := ValidateIdentifier(name); err != nil {
		return fmt.Errorf("invalid package name: %w", err)
	}

	return nil
}

func ValidateIdentifier(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("identifier cannot be empty")
	}

	if !pxIdentifier.MatchString(name) {
		return fmt.Errorf("'%s' is not a valid identifier", name)
	}

	return nil
}

package utils

import (
	"errors"
	"fmt"
	"regexp"
	"unicode"
)

var pxIdentifier = regexp.MustCompile(`(?i)^[a-z]+[a-z0-9_]*$`)

func ValidateQueryGetterSuffix(suffix string) error {
	if err := ValidateIdentifier(suffix); err != nil {
		return fmt.Errorf("invalid query getter function suffix: %w", err)
	}

	return nil
}

func ValidatePackageName(name string) error {
	if err := ValidateIdentifier(name); err != nil {
		return fmt.Errorf("invalid package name: %w", err)
	}

	return nil
}

func ValidateIdentifier(name string) error {
	if len(name) == 0 {
		return errors.New("identifier cannot be empty")
	}

	if !pxIdentifier.MatchString(name) {
		return fmt.Errorf("'%s' is not a valid identifier", name)
	}

	return nil
}

func FirstUpper(inp string) string {
	runes := []rune(inp)
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}

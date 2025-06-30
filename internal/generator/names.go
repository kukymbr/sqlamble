package generator

import (
	"strings"
	"unicode"

	"github.com/kukymbr/sqlamble/internal/utils"
)

// Cases:
//
// name
// test-name
// test_name
// test name
// TestName
// Test_Name
// Test Name
// test.name
func nameToParts(name string) []string {
	name = strings.TrimFunc(name, unicodeSepFunc)
	parts := strings.FieldsFunc(name, unicodeSepFunc)
	res := make([]string, 0, len(parts))

	for _, part := range parts {
		if err := utils.ValidateIdentifier(part); err != nil {
			continue
		}

		part = strings.ToLower(part)

		res = append(res, part)
	}

	return res
}

func unicodeSepFunc(r rune) bool {
	return r == '-' || r == '_' || r == '.' || r == ':' || unicode.IsSpace(r)
}

func partsToCapitalized(parts []string, firstLower bool) string {
	res := ""

	start := 0
	if firstLower {
		start = 1
	}

	for i := start; i < len(parts); i++ {
		part := parts[i]

		runes := []rune(part)
		runes[0] = unicode.ToTitle(runes[0])

		res += string(runes)
	}

	return res
}

func partsToIdentifier(parts []string) string {
	return strings.Join(parts, "-")
}

func firstLower(inp string) string {
	runes := []rune(inp)
	runes[0] = unicode.ToLower(runes[0])

	return string(runes)
}

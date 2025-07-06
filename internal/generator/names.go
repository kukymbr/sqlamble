package generator

import (
	"strings"
	"unicode"

	"github.com/kukymbr/sqlamble/internal/utils"
)

func nameToParts(name string) []string {
	runes := []rune(name)
	parts := make([]string, 1)
	index := 0

	next := func() {
		if err := utils.ValidateIdentifier(parts[index]); err != nil {
			parts[index] = ""

			return
		}

		parts[index] = strings.ToLower(parts[index])
		parts = append(parts, "")

		index++
	}

	for i, r := range runes {
		if unicode.IsUpper(r) {
			if len(parts[index]) > 0 &&
				// do not split UPPERCASE words
				i >= 1 && !unicode.IsUpper(runes[i-1]) {
				next()
			}

			parts[index] += string(r)

			continue
		}

		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			parts[index] += string(r)

			continue
		}

		next()
	}

	next()

	if parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}

	return parts
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

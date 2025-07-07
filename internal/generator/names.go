package generator

import (
	"strings"
	"unicode"
)

const (
	runeStateNone byte = iota
	runeStateUpper
	runeStateLower
	runeStateNumber
)

//nolint:cyclop,funlen
func nameToWords(name string) []string {
	var (
		words []string
		word  []rune
		state = runeStateNone
	)

	writeWord := func() {
		if len(word) == 0 {
			return
		}

		s := string(word)
		word = nil
		state = runeStateNone

		words = append(words, s)
	}

	addToWord := func(r rune, runeType byte) {
		if runeType == runeStateUpper {
			r = unicode.ToLower(r)
		}

		state = runeType

		word = append(word, r)
	}

	//nolint:staticcheck
	for _, r := range name {
		if unicode.IsDigit(r) {
			addToWord(r, runeStateNumber)

			continue
		}

		if !unicode.IsLetter(r) {
			writeWord()

			continue
		}

		var runeType byte

		if unicode.IsUpper(r) {
			runeType = runeStateUpper
		} else {
			runeType = runeStateLower
		}

		breakWord := state != runeStateNone
		//nolint:staticcheck
		// Ignoring QF1001 (could apply De Morgan's law) I find it more readable as it is now.
		breakWord = breakWord && !(state == runeStateUpper && len(word) == 1)
		breakWord = breakWord && (state == runeStateNumber || state != runeType)

		if breakWord {
			writeWord()
		}

		addToWord(r, runeType)
	}

	if len(word) > 0 {
		writeWord()
	}

	return words
}

func wordsToCapitalized(parts []string) string {
	res := ""
	start := 0

	for i := start; i < len(parts); i++ {
		part := parts[i]

		runes := []rune(part)
		runes[0] = unicode.ToTitle(runes[0])

		res += string(runes)
	}

	return res
}

func wordsToIdentifier(parts []string) string {
	return strings.Join(parts, "-")
}

func firstLower(inp string) string {
	runes := []rune(inp)
	runes[0] = unicode.ToLower(runes[0])

	return string(runes)
}

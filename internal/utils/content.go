package utils

const quote = '`'

func GetQuotedContent(content string) string {
	runes := []rune(content)
	res := make([]rune, 0, len(runes)+2)
	res = append(res, quote)

	for i, r := range runes {
		if r != quote {
			if i > 0 && runes[i-1] == quote {
				res = append(res, '"', '+', quote)
			}

			res = append(res, r)

			continue
		}

		if i > 0 && runes[i-1] != quote {
			res = append(res, quote, '+', '"')
		}

		res = append(res, r)
	}

	res = append(res, quote)

	return string(res)
}

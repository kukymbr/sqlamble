package utils_test

import (
	"testing"

	"github.com/kukymbr/sqlamble/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetQuotedContent(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "SELECT * FROM test",
			Expected: "`SELECT * FROM test`",
		},
		{
			Input:    "SELECT '`1`' AS data FROM test",
			Expected: "`SELECT '`+\"`\"+`1`+\"`\"+`' AS data FROM test`",
		},
		{
			Input:    "SELECT '```shell\n echo test```' AS md FROM test",
			Expected: "`SELECT '`+\"```\"+`shell\n echo test`+\"```\"+`' AS md FROM test`",
		},
	}

	for _, test := range tests {
		t.Run(test.Input, func(t *testing.T) {
			quoted := utils.GetQuotedContent(test.Input)

			assert.Equal(t, test.Expected, quoted)
		})
	}
}

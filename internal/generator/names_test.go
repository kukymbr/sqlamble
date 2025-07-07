package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameToParts(t *testing.T) {
	tests := []struct {
		Input    string
		Expected []string
	}{
		{Input: "name", Expected: []string{"name"}},
		{Input: "test-name", Expected: []string{"test", "name"}},
		{Input: "test_name", Expected: []string{"test", "name"}},
		{Input: "test name", Expected: []string{"test", "name"}},
		{Input: "TestName", Expected: []string{"test", "name"}},
		{Input: "Test_Name", Expected: []string{"test", "name"}},
		{Input: "Test Name", Expected: []string{"test", "name"}},
		{Input: "test.name", Expected: []string{"test", "name"}},
		{Input: "test, name", Expected: []string{"test", "name"}},
		{Input: "test -/- name", Expected: []string{"test", "name"}},
		{Input: "test name.", Expected: []string{"test", "name"}},
		{Input: "testNAME", Expected: []string{"test", "name"}},
		{Input: "__TEST_name.", Expected: []string{"test", "name"}},
		{Input: "TESTname.", Expected: []string{"test", "name"}},
		{Input: "TEST1name", Expected: []string{"test1", "name"}},
		{Input: "TEST123name", Expected: []string{"test123", "name"}},
		{Input: "1TESTname", Expected: []string{"1", "test", "name"}},
	}

	for _, test := range tests {
		t.Run(test.Input, func(t *testing.T) {
			parts := nameToWords(test.Input)

			assert.Equal(t, test.Expected, parts)
		})
	}
}

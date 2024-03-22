// Тестируем функцию AnswerTrim
package test

import (
	"pet_calculator/pkg/decor"
	"testing"
)

func TestAnswerTrim(t *testing.T) {

	testTable := []struct {
		result   float64
		expected string
	}{
		{
			result:   54.33,
			expected: "54.33",
		},

		{
			result:   0,
			expected: "0",
		},

		{
			result:   54.1000,
			expected: "54.1",
		},
	}

	for _, i := range testTable {
		actual := decor.AnswerTrim(i.result)

		if i.expected != actual {
			t.Errorf("Result was incorrect, got: %s, want: %s.", actual, i.expected)
		}

	}

}

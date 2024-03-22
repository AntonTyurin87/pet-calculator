package test

import (
	"pet_calculator/pkg/decor"
	"testing"
)

func TestDecorPattern(t *testing.T) {

	testTable := []struct {
		DecorElemetn string
		DecoreLen    int
		expected     string
	}{
		{
			DecorElemetn: "*",
			DecoreLen:    1,
			expected:     "*",
		},

		{
			DecorElemetn: "+++",
			DecoreLen:    2,
			expected:     "++++++",
		},

		{
			DecorElemetn: "+++",
			DecoreLen:    -1,
			expected:     "",
		},
	}

	for _, i := range testTable {
		actual := decor.DecorPattern(i.DecorElemetn, i.DecoreLen)

		if i.expected != actual {
			t.Errorf("Result was incorrect, got: %s, want: %s.", actual, i.expected)
		}

	}

}

package parse

import (
	"math"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input string
		value float64
	}{
		{"1 + 2 * 3", 7},
		{"7 - 9 * (2 - 3)", 16},
		{"2 * 3 * 4", 24},
		{"2 ^ 3 ^ 4", math.Pow(2, math.Pow(3, 4))},
		{"(2 ^ 3) ^ 4", 4096},
		{"5", 5},
		{"4 + 2", 6},
		{"9 - 8 - 7", -6},
		{"9 - (8 - 7)", 8},
		{"9 - 8) - 7", -6},
		{"2 + 3 ^ 2 * 3 + 4", 33},
	}

	p := NewParser()
	tolerance := 0.00000000001

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			t.Parallel()

			expr, err := p.Parse(test.input)

			if err != nil {
				t.Fatal(err)
			}

			expected := test.value
			actual := expr.Eval()

			if math.Abs(expected-actual) > tolerance {
				t.Errorf("Expected %f, got %f", expected, actual)
			}
		})
	}
}
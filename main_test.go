package main_test

import (
	"testing"

	"github.com/paveldroo/advent-of-code-2024/day3"
)

func TestMain(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			expected: 161,
		},
	}

	for _, tt := range tests {
		res := day3.First(tt.input)
		if tt.expected != res {
			t.Fatalf("res: %v, expected: %v", res, tt.expected)
		}
	}
}

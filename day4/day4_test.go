package day4_test

import (
	"testing"

	"github.com/paveldroo/advent-of-code-2024/day4"
)

func TestRun_1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX",
			expected: 18,
		},
	}

	for _, tt := range tests {
		res := day4.Run_1(tt.input)
		if res != tt.expected {
			t.Fatalf("input: %v, expected: %v, res: %v", tt.input, tt.expected, res)
		}
	}
}

func TestRun_2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX",
			expected: 9,
		},
	}

	for _, tt := range tests {
		res := day4.Run_2(tt.input)
		if res != tt.expected {
			t.Fatalf("input: %v, expected: %v, res: %v", tt.input, tt.expected, res)
		}
	}
}

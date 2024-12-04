package day3_test

import (
	"reflect"
	"testing"

	"github.com/paveldroo/advent-of-code-2024/day3"
)

func TestProcessInts(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
		wantErr  bool
	}{
		{
			input:    "(1,5)",
			expected: []int{1, 5},
			wantErr:  false,
		},
		{
			input:    "(1555,55555)",
			expected: []int{1555, 55555},
			wantErr:  false,
		},
		{
			input:    "(15)",
			expected: []int{15},
			wantErr:  true,
		},
		{
			input:    "(1,)",
			expected: []int{1},
			wantErr:  true,
		},
		{
			input:    "(,5)",
			expected: []int{},
			wantErr:  true,
		},
		{
			input:    "(&)",
			expected: []int{},
			wantErr:  true,
		},
		{
			input:    "(1,5&",
			expected: []int{1},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		ints, err := day3.ProcessInts([]rune(tt.input))
		if err != nil && !tt.wantErr {
			t.Fatalf("input: %v, got error: %v", tt.input, err)
		}
		if !reflect.DeepEqual(tt.expected, ints) {
			t.Fatalf("input: %v, ints: %v, expected: %v", tt.input, ints, tt.expected)
		}
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
		wantOk   bool
	}{
		{
			input:    "mul(15,55)",
			expected: []int{15, 55},
			wantOk:   true,
		},
		{
			input:    "mul(15)",
			expected: []int(nil),
			wantOk:   false,
		},
		{
			input:    "mu&(2,4)",
			expected: []int(nil),
			wantOk:   false,
		},
	}

	for _, tt := range tests {
		ints, ok := day3.Process([]rune(tt.input))
		if !ok && tt.wantOk {
			t.Fatalf("input: %v, expected: %v, ints: %v, got not ok", tt.input, tt.expected, ints)
		}
		if !reflect.DeepEqual(tt.expected, ints) {
			t.Fatalf("input: %v, ints: %v, expected: %v", tt.input, ints, tt.expected)
		}
	}
}

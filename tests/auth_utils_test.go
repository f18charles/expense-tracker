package test

import (
	"github.com/f18charles/expense-tracker/internal/utils"
	"testing"
)

func TestAutoIncrementID(t *testing.T) {
	type TestCase struct {
		input     int
		want int
	}

	tests := []TestCase{
		{input: 1, want: 2},
		{input: 0, want: 1},
		{input: 5, want: 6},
	}

	for _, test := range tests {
		got, err := utils.AutoIncrementinput(test.input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got != test.want {
			t.Errorf("AutoIncrementinput(%d) = %d, want %d", test.input, got, test.want)
		}
	}
}

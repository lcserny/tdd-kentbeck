package tdd_kentbeck

import "testing"

func TestMultiplication(t *testing.T) {
	fiveDollars := NewDollar(5)
	fiveDollars.Times(2)

	if fiveDollars.Amount() != 10 {
		t.Errorf("amounts not equal, got %d, want %d", fiveDollars.Amount(), 10)
	}
}

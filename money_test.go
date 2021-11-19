package tdd_kentbeck

import "testing"

func TestMultiplication(t *testing.T) {
	fiveDollars := NewDollar(5)
	tenDollars := fiveDollars.Times(2)

	want := NewDollar(10)
	if tenDollars != want {
		t.Errorf("amounts not equal, got %d, want %d", tenDollars, want)
	}

	fifteenDollars := fiveDollars.Times(3)

	want = NewDollar(15)
	if fifteenDollars != want {
		t.Errorf("amounts not equal, got %d, want %d", fifteenDollars, want)
	}

	fiveFrancs := NewFranc(5)
	tenFrancs := fiveFrancs.Times(2)

	wantFranc := NewFranc(10)
	if tenFrancs != wantFranc {
		t.Errorf("amounts not equal, got %d, want %d", tenFrancs, wantFranc)
	}

	fifteenFrancs := fiveFrancs.Times(3)

	wantFranc = NewFranc(15)
	if fifteenFrancs != wantFranc {
		t.Errorf("amounts not equal, got %d, want %d", fifteenFrancs, wantFranc)
	}
}

func TestEquality(t *testing.T) {
	fiveDollar := NewDollar(5)
	newFiveDollar := NewDollar(5)

	if fiveDollar != newFiveDollar {
		t.Errorf("two 5 dollar bills are not equal, first %d, second %d", fiveDollar, newFiveDollar)
	}

	sixDollar := NewDollar(6)
	sevenDollar := NewDollar(7)

	if sixDollar == sevenDollar {
		t.Errorf("different dollar bills should not be equal, first %d, second %d", sixDollar, sevenDollar)
	}

	fiveFranc := NewFranc(5)
	newFiveFranc := NewFranc(5)

	if fiveFranc != newFiveFranc {
		t.Errorf("two 5 Franc bills are not equal, first %d, second %d", fiveFranc, newFiveFranc)
	}

	sixFranc := NewFranc(6)
	sevenFranc := NewFranc(7)

	if sixFranc == sevenFranc {
		t.Errorf("different Franc bills should not be equal, first %d, second %d", sixFranc, sevenFranc)
	}
}

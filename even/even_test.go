package even

import (
	"testing"
)

func TestEven(t *testing.T) {
		if got, want := IsEven(6), true; got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestNotEven(t *testing.T) {
		if got, want := IsEven(7), false; got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}
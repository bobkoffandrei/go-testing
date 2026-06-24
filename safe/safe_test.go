package safe

import (
	"testing"
)

func TestSafe(t *testing.T) {
	if got, want := MustAt([]int{10, 20, 30}, 1), 20; want != got {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMustAt_Panic_OutOfRange(t *testing.T) {
	xs := []int{1, 2}
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic, got none")
		}
	}()
	_ = MustAt(xs, 2)
}

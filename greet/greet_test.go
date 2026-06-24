package greet

import (
	"testing"
	//"errors"
)



func TestHello(t *testing.T) {
		cases := []struct {
		g string
		want string
	}{
		{"Go", "Hello, Go"},
		{"World", "Hello, World"},
		{"Гофер", "Hello, Гофер"},
		{" Go ", "Hello,  Go "},

	}

	for _, n := range cases {
	got, _ := Hello(n.g)
	assertString(t, got, n.want)
	}

	_, err := Hello("")
	assertError(t, err, true)

}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	 
		if want != got {
		t.Fatalf("got: %v, want: %v", got, want)
	}
}

func assertError(t testing.TB, got error, want bool) {
	t.Helper()
	if got == nil {
		t.Fatal("ожидали ошибку, но получили nil")
	}
	if got.Error() != "name cannot be empty" {
		t.Errorf("получили ошибку %v, хотели %v", got.Error(), want)
	}
}
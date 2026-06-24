package normalize

import (
	"fmt"
	"testing"
)

func TestClean(t *testing.T) {
		cases := []struct {
		s string
		want string
	}{
		{"  Hello", "hello"},
		{"  GO			Lang", "go lang"},
		{"GOLANG", "golang"},
		{"", ""},
		{"./././.", "./././."},

	}

	for _, c := range cases {
		name := fmt.Sprintf("%s", c.s)

			t.Run(name, func(t *testing.T) {
			got := Clean(c.s)
			if got != c.want {
				t.Errorf("Clean(%s) = %s, хотели %s", c.s, got, c.want)
			}
		})
	}



}
	

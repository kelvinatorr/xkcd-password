package main

import (
	"io"
	"testing"
)

type MockReader struct {
	data string
}

func (r MockReader) String() string {
	return r.data
}

func (r MockReader) Read(b []byte) (int, error) {
	for i, val := range []byte(r.data) {
		b[i] = val
	}
	return len(b), io.EOF
}

func TestLineCounter(t *testing.T) {
	var tests = []struct {
		input MockReader
		want  int
	}{
		{MockReader{"cow\n level\n"}, 2},
		{MockReader{"cow"}, 0},
		{MockReader{"cow\n"}, 1},
	}

	for _, test := range tests {
		got, err := lineCounter(test.input)
		if err != nil {
			t.Errorf("lineCounter(%q) returned an error", test.input)
		}

		if got != test.want {
			t.Errorf("lineCounter(%q) = %d", test.input, got)
		}
	}

}

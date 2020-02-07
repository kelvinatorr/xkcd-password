package main

import (
	"io"
	"strings"
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

func TestGeneratePassword(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{4, 4},
		{4751, 4750},
		{23, 23},
		{8, 8},
		{4750, 4750},
	}
	for _, test := range tests {
		got := GeneratePassword(test.input)
		// Test that we get the number of words we requested.
		gotWordSlice := strings.Split(got, " ")
		if len(gotWordSlice) != test.want {
			t.Errorf("GeneratePassword(%d) got %d words; want %d", test.input, len(gotWordSlice), test.want)
		}
		// Test that none of the words repeat.
		wordMap := make(map[string]bool)
		for _, word := range gotWordSlice {
			if _, present := wordMap[word]; present {
				t.Errorf("GeneratePassword(%d) repeats word: %v", test.input, word)
			} else {
				wordMap[word] = true
			}
		}
	}
}

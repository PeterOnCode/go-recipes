package bytestrings

import (
	"bytes"
	"errors"
	"testing"
)

func TestBuffer(t *testing.T) {
	var bufferTests = []struct {
		input  string
		output *bytes.Buffer
	}{
		{"abc", bytes.NewBufferString("abc")},
		{"", bytes.NewBufferString("")},
		{"❤️", bytes.NewBufferString("❤️")},
	}
	for _, bt := range bufferTests {
		if got, want := Buffer(bt.input), bt.output; got.String() != want.String() {
			t.Errorf("got: %#v, want: %#v", got, want)
		}
	}
}

func TestToString(t *testing.T) {
	var toStringTests = []struct {
		input   *bytes.Buffer
		output1 string
		output2 error
	}{
		{bytes.NewBufferString("abc"), "abc", nil},
	}

	for _, tst := range toStringTests {
		got1, got2 := toString(tst.input)
		if got1 != tst.output1 || !errors.Is(got2, tst.output2) {
			t.Errorf("got: %#v %#v, want: %#v %#v", got1, got2, tst.output1, tst.output2)
		}
	}
}

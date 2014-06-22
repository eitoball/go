package secret

import (
	"reflect"
	"testing"
)

var tests = []struct {
	code int
	h    []string
}{
	{1, []string{"wink"}},
	{2, []string{"double blink"}},
	{4, []string{"close your eyes"}},
	{8, []string{"jump"}},
	{3, []string{"wink", "double blink"}},
	{19, []string{"double blink", "wink"}},
	{31, []string{"jump", "close your eyes", "double blink", "wink"}},
	{0, nil},
	{-1, nil},
	{32, nil},
}

func TestHandshake(t *testing.T) {
	for _, test := range tests {
		h := Handshake(test.code)
		if len(test.h) == 0 && len(h) != 0 || !reflect.DeepEqual(h, test.h) {
			t.Fatalf("Handshake(%d) = %v, want %v.", test.code, h, test.h)
		}
	}
}

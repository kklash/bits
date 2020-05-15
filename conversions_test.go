package bits

import (
	"testing"
)

func TestBits_String(t *testing.T) {
	test := func(bits Bits, expected string) {
		if actual := bits.String(); actual != expected {
			t.Errorf("Failed to generate expected bit string\nWanted %s\nGot    %s", expected, actual)
		}
	}

	test(Bits{false, true, false, true}, "0101")
	test(Bits{false, false, false, true}, "0001")
	test(Bits{}, "")
	test(Bits{true}, "1")
}

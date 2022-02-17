package bits

import (
	"testing"
)

func TestBits_Trim(t *testing.T) {
	test := func(bitString, expectedBits string) {
		bits, _ := StringToBits(bitString)
		bits = bits.Trim()
		if bits.String() != expectedBits {
			t.Errorf("Failed to trim bits\nWanted %s\nGot    %s", expectedBits, bits)
			return
		}
	}

	test("000100010101", "100010101")
	test("00000000", "")
	test("10000000000000000", "10000000000000000")
	test("00000000000000001", "1")
}

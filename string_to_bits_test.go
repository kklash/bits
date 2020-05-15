package bits

import (
	"errors"
	"testing"
)

func TestStringToBits(t *testing.T) {
	test := func(input string, valid bool) {
		bits, err := StringToBits(input)
		if valid && err != nil {
			t.Errorf("Expected error to be nil for valid input")
			return
		} else if !valid {
			if !errors.Is(err, ErrBadBitString) {
				t.Errorf("Incorrect error for invalid input\nWanted %s\nGot    %s", ErrBadBitString, err)
			}
			return
		} else if len(bits) != len(input) {
			t.Errorf("Expected input and output to have same length\nWanted %d\nGot    %d", len(input), len(bits))
			return
		}

		for i, bit := range bits {
			if bit != (input[i] == '1') {
				t.Errorf("Input-output mismatch converting from string at index %d: %s", i, input)
				return
			}
		}
	}

	test("100102003000", false)
	test("ab", false)
	test(" ", false)
	test("1010 100", false)
	test("9", false)

	test("", true)
	test("100", true)
	test("000", true)
	test("010101010010101010", true)
	test("1", true)
	test("0", true)
}

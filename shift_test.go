package bits

import (
	"fmt"
	"testing"
)

func TestBits_ShiftRight(t *testing.T) {
	test := func(input string, nShift uint32, expected string) {
		bits, _ := StringToBits(input)
		actual := bits.ShiftRight(nShift).String()
		if actual != expected {
			t.Errorf("Failed to get expected right-shifted bits\nWanted %s\nGot    %s", expected, actual)
		}
	}

	test("100", 1, "10")
	test("10100", 0, "10100")
	test("10100", 1, "1010")
	test("0010100", 1, "001010")
	test("10100", 2, "101")
	test("1", 1, "")
	test("1", 0, "1")
	test("1", 4, "")
	test("", 1, "")
}

func ExampleBits_ShiftRight() {
	bits, _ := StringToBits("11111111")
	afterShift := bits.ShiftRight(4)
	fmt.Printf("before: %s\n", bits)
	fmt.Printf("after:  %s\n", afterShift)

	// output:
	// before: 11111111
	// after:  1111
}

func TestBits_ShiftLeft(t *testing.T) {
	test := func(input string, nShift uint32, expected string) {
		bits, _ := StringToBits(input)
		actual := bits.ShiftLeft(nShift).String()
		if actual != expected {
			t.Errorf("Failed to get expected left-shifted bits\nWanted %s\nGot    %s", expected, actual)
		}
	}

	test("100", 1, "1000")
	test("10100", 0, "10100")
	test("10100", 1, "101000")
	test("0010100", 1, "00101000")
	test("10100", 2, "1010000")
	test("1", 1, "10")
	test("1", 0, "1")
	test("1", 4, "10000")
	test("", 4, "0000")
}

func ExampleBits_ShiftLeft() {
	bits, _ := StringToBits("11111111")
	afterShift := bits.ShiftLeft(4)
	fmt.Printf("before: %s\n", bits)
	fmt.Printf("after:  %s\n", afterShift)

	// output:
	// before: 11111111
	// after:  111111110000
}

package bits

import (
	"fmt"
	"testing"
)

func TestBits_PadLeft(t *testing.T) {
	test := func(n uint32, bitString, expectedBits string) {
		bits, _ := StringToBits(bitString)
		if bits == nil {
			t.Errorf("Failed to parse bits from string: %s", bitString)
			return
		}

		bits = bits.PadLeft(n)
		if bits.String() != expectedBits {
			t.Errorf("Failed to generate expected bits\nWanted %s\nGot    %s", expectedBits, bits)
		}
	}

	test(8, "00001", "00000001")
	test(11, "01010101", "00001010101")
	test(0, "11110000", "11110000")
	test(1, "10010", "10010")
	test(2, "1000001", "01000001")
	test(5, "", "00000")
}

func ExampleBits_PadRight() {
	// Pad a bit slice to an even multiple of some number.
	// Makes it easier to perform loops on a bit slice.
	bits, _ := StringToBits("111100001111")
	bits = bits.PadRight(8)
	fmt.Println(bits)
	// output:
	// 1111000011110000
}

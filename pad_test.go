package bits

import (
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

func TestSplitPadLeft(t *testing.T) {
	test := func(n uint32, bitString string, expectedBitGroups ...string) {
		bits, _ := StringToBits(bitString)
		if bits == nil {
			t.Errorf("Failed to parse bits from string: %s", bitString)
			return
		}

		bitGroups := bits.SplitPadLeft(n)
		if len(bitGroups) != len(expectedBitGroups) {
			t.Errorf("Split bit groups length does not match\nWanted %d\nGot    %d", len(expectedBitGroups), len(bitGroups))
			return
		}

		for i := 0; i < len(bitGroups); i++ {
			if group := bitGroups[i].String(); group != expectedBitGroups[i] {
				t.Errorf("Split bit group does not match\nWanted %s\nGot    %s", expectedBitGroups[i], group)
			}
		}
	}

	test(8, "101010101111111100000000", "10101010", "11111111", "00000000")
	test(8, "111100001111", "00001111", "00001111")
	test(11, "101010101010101010", "00001010101", "01010101010")
	test(1, "00001111", "0", "0", "0", "0", "1", "1", "1", "1")
	test(0, "00001111")
}

package bits

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBits_Split(t *testing.T) {
	test := func(n uint32, input string, expected ...string) {
		expectedGroups := make([]Bits, len(expected))
		for i, group := range expected {
			expectedGroups[i], _ = StringToBits(group)
		}

		inputBits, _ := StringToBits(input)
		groups := inputBits.Split(n)

		if !reflect.DeepEqual(groups, expectedGroups) {
			t.Errorf("split groups do not match\nWanted %s\nGot    %s", expectedGroups, groups)
			return
		}
	}

	test(4, "00001111", "0000", "1111")
	test(1, "1001", "1", "0", "0", "1")
	test(1, "1001", "1", "0", "0", "1")
	test(8, "11111", "11111")
	test(3, "11111", "111", "11")
	test(2, "11111", "11", "11", "1")
	test(4, "")
	// TODO test .Split(0) case
}

func TestBits_SplitPadLeft(t *testing.T) {
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

func ExampleBits_Split() {
	bitSlice := BytesToBits([]byte{1, 2, 3, 4, 5})
	fmt.Println(bitSlice.Split(8))
	fmt.Println(bitSlice.Split(5))

	// Output:
	// [00000001 00000010 00000011 00000100 00000101]
	// [00000 00100 00001 00000 00110 00001 00000 00101]
}

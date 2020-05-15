package bits

import (
	"testing"
)

func TestUintToBits(t *testing.T) {
	test := func(inputNum interface{}, expectedBits string) {
		bits := UintToBits(inputNum)
		if bits.String() != expectedBits {
			t.Errorf("Failed to parse uint as bits\nWanted %s\nGot    %s", expectedBits, bits)
			return
		}
	}

	test(uint8(8), "00001000")
	test(byte(3), "00000011")
	test(uint64(10), "0000000000000000000000000000000000000000000000000000000000001010")
	test(uint32(5883), "00000000000000000001011011111011")
}

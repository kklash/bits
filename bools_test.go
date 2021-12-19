package bits

import (
	"reflect"
	"testing"
)

func TestBits_Bools(t *testing.T) {
	test := func(bitString string) {
		bits, _ := StringToBits(bitString)

		bools := bits.Bools()

		for i := 0; i < len(bitString); i++ {
			if bool(bits[i]) != bools[i] {
				t.Errorf("expected bools and Bits to match")
				return
			}
		}

		bitsFromBools := BoolsToBits(bools)

		if !reflect.DeepEqual(bits, bitsFromBools) {
			t.Errorf("expected BoolsToBits to parse bools as bits correctly")
			return
		}
	}

	test("")
	test("1")
	test("01")
	test("101")
	test("0001")
	test("01001")
}

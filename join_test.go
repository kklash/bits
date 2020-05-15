package bits

import (
	"testing"
)

func TestJoin(t *testing.T) {
	test := func(expectedBits string, groups ...string) {
		bitGroups := make([]Bits, len(groups))
		for i := 0; i < len(groups); i++ {
			bitGroups[i], _ = StringToBits(groups[i])
		}

		joined := Join(bitGroups)
		if joined.String() != expectedBits {
			t.Errorf("Failed to join bits\nExpected %s\nGot    %s", expectedBits, joined)
			return
		}
	}

	test("001011100101010001110", "001", "011100101", "010001110")
	test("0000111100001111", "00001111", "00001111")
	test("000000", "000", "nil", "000")
	test("")
	test("0111111111111", "0", "111111111111")
}

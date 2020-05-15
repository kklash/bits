package bits

import (
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

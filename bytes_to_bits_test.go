package bits

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestByteToBits(t *testing.T) {
	test := func(b byte, expectedBits string) {
		bits := ByteToBits(b)
		if bits.String() != expectedBits {
			t.Errorf("Failed to generate expected bits\nWanted %s\nGot    %s", expectedBits, bits)
		} else if fmt.Sprintf("%x", bits.Bytes()) != fmt.Sprintf("%x", []byte{b}) {
			t.Errorf("Failed to encode bits to byte\nWanted %x\nGot    %x", []byte{b}, bits.Bytes())
		}
	}

	test(0xff, "11111111")
	test(0x01, "00000001")
	test(0x00, "00000000")
	test(0x80, "10000000")
	test(0xf0, "11110000")
	test(0xaa, "10101010")
}

func TestBytesToBits(t *testing.T) {
	test := func(h, expectedBits string) {
		bytes, _ := hex.DecodeString(h)
		bits := BytesToBits(bytes)
		if bits.String() != expectedBits {
			t.Errorf("Failed to generate expected bits\nWanted %s\nGot    %s", expectedBits, bits)
			return
		}
	}

	test("ffaa", "1111111110101010")
	test("ccee", "1100110011101110")
	test("1122334455", "0001000100100010001100110100010001010101") // Tests for stripping of leading zeros
	test("ddddaaaa01", "1101110111011101101010101010101000000001")
}

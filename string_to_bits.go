package bits

import (
	"errors"
)

// ErrBadBitString is returned when parsing an invalid bit string.
var ErrBadBitString = errors.New("Failed to parse string as Bits, every character must be '1' or '0'")

// StringToBits attempts to parse a string as a slice of Bits, returning ErrBadBitString on failure.
func StringToBits(bitString string) (Bits, error) {
	bits := make(Bits, len(bitString))
	for i, s := range bitString {
		if s != '1' && s != '0' {
			return nil, ErrBadBitString
		}

		bits[i] = s == '1'
	}

	return bits, nil
}

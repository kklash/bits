package bits

// ByteToBits converts a byte into a Bit slice of length 8.
func ByteToBits(b byte) Bits {
	bits := make(Bits, 8)
	for i := 7; i >= 0; i-- {
		bits[i] = b&1 == 1
		b >>= 1
	}

	return bits
}

// BytesToBits parses a byte slice and converts it into a Bit slice. The length of
// the returned Bit slice will always be len(bytes) * 8.
func BytesToBits(bytes []byte) Bits {
	if bytes == nil {
		return nil
	}

	bits := make(Bits, len(bytes)*8)
	for i := 0; i < len(bytes); i++ {
		bitIndex := i * 8
		newBits := ByteToBits(bytes[i])
		copy(bits[bitIndex:], newBits)
	}

	return bits
}

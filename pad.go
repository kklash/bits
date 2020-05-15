package bits

// PadLeft pads the Bit slice with leading zeros until len(bits) % n == 0 and returns the result.
func (bits Bits) PadLeft(n uint32) Bits {
	var remainder uint32
	if n == 0 {
		return bits
	} else if remainder = bits.BitSize() % n; remainder == 0 && bits.BitSize() != 0 {
		return bits
	}

	// Pad leading zeros
	padding := make(Bits, int(n-remainder))
	return append(padding, bits...)
}

// PadRight pads the Bit slice with trailing zeros until len(bits) % n == 0 and returns the result.
// If len(bits) == 0, it pads until len(bits) == n
func (bits Bits) PadRight(n uint32) Bits {
	var remainder uint32
	if n == 0 {
		return bits
	} else if remainder = bits.BitSize() % n; remainder == 0 && bits.BitSize() != 0 {
		return bits
	}

	// Pad trailing zeros
	return bits.ShiftLeft(n - remainder)
}

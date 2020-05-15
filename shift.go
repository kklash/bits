package bits

// ShiftLeft shifts bits by n bits to the left (appends zeros) and returns the result.
// Note that if there are leading zeros at the beginning of bits, they are not affected.
func (bits Bits) ShiftLeft(n uint32) Bits {
	endBits := make(Bits, int(n))
	return append(bits, endBits...)
}

// ShiftRight shifts bits by n bits to the right (cuts off trailing bits) and returns
// the result. Note that if there are leading zeros at the beginning of bits, they are
// not affected. If n >= len(bits), returns an empty Bit slice.
func (bits Bits) ShiftRight(n uint32) Bits {
	if n >= bits.BitSize() {
		return Bits{}
	}

	return bits[:bits.BitSize()-n]
}

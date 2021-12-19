package bits

// Bools returns the Bits as a slice of bools.
func (bits Bits) Bools() []bool {
	bools := make([]bool, len(bits))
	for i, bit := range bits {
		bools[i] = bool(bit)
	}
	return bools
}

// BoolsToBits converts the given slice of bools into an array of Bits.
func BoolsToBits(bools []bool) Bits {
	bits := make(Bits, len(bools))
	for i, b := range bools {
		bits[i] = Bit(b)
	}
	return bits
}

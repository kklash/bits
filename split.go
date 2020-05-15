package bits

// Split splits a Bit slice into groups of n bits. If len(bits) % n != 0, then
// the last group in the slice will be of length len(bits) % n (the remainder).
func (bits Bits) Split(n uint32) []Bits {
	if n == 0 {
		return nil
	}

	nGroups := len(bits) / int(n)
	nLeftover := len(bits) % int(n)
	if nLeftover != 0 {
		// Extra group for leftovers
		nGroups++
	}

	bitGroups := make([]Bits, nGroups)

	for i := 0; i < len(bitGroups); i++ {
		var groupSize int
		if i == len(bits)/int(n) {
			groupSize = nLeftover
		} else {
			groupSize = int(n)
		}

		bitGroups[i] = make(Bits, groupSize)

		for j := 0; j < groupSize; j++ {
			bitIndex := i*int(n) + j
			bitGroups[i][j] = bits[bitIndex]
		}
	}

	return bitGroups
}

// SplitPadLeft splits a Bit slice into groups of n bits. If len(bits) % n != 0,
// bits is padded with leading zeros until len(bits) % n == 0, before being split.
func (bits Bits) SplitPadLeft(n uint32) []Bits {
	if n == 0 {
		return nil
	}

	bits = bits.PadLeft(n)
	return bits.Split(n)
}

// SplitPadLeft splits a Bit slice into groups of n bits. If len(bits) % n != 0,
// bits is padded with trailing zeros until len(bits) % n == 0, before being split.
func (bits Bits) SplitPadRight(n uint32) []Bits {
	if n == 0 {
		return nil
	}

	bits = bits.PadRight(n)
	return bits.Split(n)
}

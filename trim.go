package bits

// LeadingZeros determines the number of zeros leading the
// bits (from the left) and returns the result.
func (bits Bits) LeadingZeros() uint32 {
	var leadingZeros int
	for leadingZeros = 0; leadingZeros < len(bits) && !bits[leadingZeros]; leadingZeros++ {
	}
	return uint32(leadingZeros)
}

// Trim trims all leading zeros from bits and returns the result.
func (bits Bits) Trim() Bits {
	return bits[bits.LeadingZeros():]
}

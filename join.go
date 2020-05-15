package bits

// Join concatenates the Bits inside a slice of Bit slices.
func Join(bitGroups []Bits) Bits {
	if bitGroups == nil {
		return nil
	}

	totalSize := 0
	for i := 0; i < len(bitGroups); i++ {
		if bitGroups[i] != nil {
			totalSize += len(bitGroups[i])
		}
	}

	bits := make(Bits, totalSize)

	var i int = 0
	for _, group := range bitGroups {
		if group != nil {
			for _, bit := range group {
				bits[i] = bit
				i++
			}
		}
	}

	return bits
}

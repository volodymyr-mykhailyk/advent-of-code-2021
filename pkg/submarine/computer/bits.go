package computer

func UpdateBit(n int, pos int, value int) int {
	if value == 1 {
		return SetBit(n, pos)
	} else {
		return ClearBit(n, pos)
	}
}

func SetBit(n int, pos int) int {
	n |= 1 << pos
	return n
}

func ClearBit(n int, pos int) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}

func HasBit(n int, pos int) bool {
	val := n & (1 << pos)
	return val > 0
}


package ids

// Xorshift96Rand.
func Xorshift96Rand() uint64 { //period 2^96-1
	x := uint64(123456789)
	y := uint64(362436069)
	z := uint64(521288629)
	var t uint64
	x ^= x << 16
	x ^= x >> 5
	x ^= x << 1

	t = x
	x = y
	y = z
	z = t ^ x ^ y

	return z
}

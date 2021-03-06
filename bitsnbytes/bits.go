package bitsnbytes

import (
	"fmt"
)

// CHECK_BIT_U32 checks whether a bit is set in val at pos.
func CHECK_BIT_U32(val uint32, pos uint64) uint32 {
	return ((val) & (1 << (pos)))
}

// CHECK_BIT checks whether a bit is set in val at pos.
func CHECK_BIT(val uint64, pos uint64) uint64 {
	return ((val) & (1 << (pos)))
}

// CHECK_IS_BIT checks whether a bit is set in val at pos in which case it returns true.
func CHECK_IS_BIT(val uint64, pos uint64) bool {
	return CHECK_BIT(val, pos) > 0
}

// SET_BIT sets a bit in val.
func SET_BIT(val uint64, pos uint64) uint64 {
	(val) |= (1 << pos)
	return val
}

// SET_IS_BIT sets a bit if is is true, else it clears the bit in val at pos.
func SET_IS_BIT(val uint64, pos uint64, is bool) uint64 {
	if is {
		return SET_BIT(val, pos)
	}
	return CLEAR_BIT(val, pos)
}

// CLEAR_BIT clears the bit in val at pos.
func CLEAR_BIT(val uint64, pos uint64) uint64 {
	(val) &= ^(1 << pos)
	return val
}

// BITMASK_FOR creates a bitmask for len starting at offset offs.
func BITMASK_FOR(len, offs uint64) uint64 {
	val := (int64)(^(^0 << len) << offs)
	return (uint64)(val)
}

// BitsetString returns a string representing bits in bs.
func BitsetString(bs uint64) string {
	var res string
	for i := uint64(0); i < 8; i++ {
		res += fmt.Sprintln("bitset at", i, CHECK_BIT(bs, i))
	}
	return res
}

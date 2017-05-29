package bitsnbytes

import (
	"fmt"
)

func CHECK_BIT_U32(val uint32, pos uint64) uint32 {
	return ((val) & (1 << (pos)))
}

func CHECK_BIT(val uint64, pos uint64) uint64 {
	return ((val) & (1 << (pos)))
}

func CHECK_IS_BIT(val uint64, pos uint64) bool {
	return CHECK_BIT(val, pos) > 0
}

func SET_BIT(val uint64, pos uint64) uint64 {
	(val) |= (1 << pos)
	return val
}

func SET_IS_BIT(val uint64, pos uint64, is bool) uint64 {
	if is {
		return SET_BIT(val, pos)
	}
	return CLEAR_BIT(val, pos)
}

func CLEAR_BIT(val uint64, pos uint64) uint64 {
	(val) &= ^(1 << pos)
	return val
}

func BITMASK_FOR(len, offs uint64) uint64 {
	val := (int64)(^(^0 << len) << offs)
	return (uint64)(val)
}

func BitsetString(bs uint64) string {
	var res string
	for i := uint64(0); i < 8; i++ {
		res += fmt.Sprintln("bitset at", i, CHECK_BIT(bs, i))
	}
	return res
}

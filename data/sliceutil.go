package data

// CopyBuf uses make and copy to create a new byte slice out
// with the exact size of in.
func CopyBuf(in []byte) (out []byte) {
	out = make([]byte, len(in))
	copy(out, in)
	return
}

// CopyInt64s uses make and copy to create a new int64 slice out
// with the exact size of in.
func CopyInt64s(in []int64) (out []int64) {
	out = make([]int64, len(in))
	copy(out, in)
	return
}

// CopyStrings uses make and copy to create a new string slice out
// with the exact size of in.
func CopyStrings(in []string) (out []string) {
	out = make([]string, len(in))
	copy(out, in)
	return
}

func ReverseByteSlices(vals [][]byte) {
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		vals[i], vals[j] = vals[j], vals[i]
	}
}

func Uints64Contains(vs []uint64, i uint64) bool {
	for _, v := range vs {
		if v == i {
			return true
		}
	}
	return false
}

package data

// all slices have the same size and are just "splitted"
func LargeBytesSliceFromData(data []byte, cnt, l int) (lbs [][]byte) {
	lbs = make([][]byte, cnt)
	for i := range lbs {
		beg := i * l
		lbs[i] = data[beg : beg+l]
	}
	return
}

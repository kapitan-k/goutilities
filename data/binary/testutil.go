package binary

import (
	"math/rand"
)

func RandomDatasWithRandomElemSizePrefixedMinSize(minSize, cnt uint64) [][]byte {
	datas := make([][]byte, cnt)
	for i := range datas {
		sz := uint64(rand.Int63n(1024) + ArrayLenByteSz)
		if sz < minSize {
			sz = minSize
		}
		data := make([]byte, sz)
		rand.Read(data)
		SetArrayLenAt(data, 0, uint32(len(data)-ArrayLenByteSz))
		datas[i] = data
	}

	return datas
}

func RandomDatasWithRandomElemSizePrefixed(cnt uint64) [][]byte {
	return RandomDatasWithRandomElemSizePrefixedMinSize(0, cnt)
}

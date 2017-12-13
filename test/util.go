package test

import (
	"math/rand"
)

// RandomDatas returns cnt byte slices with len of elemSz containing random data.
func RandomDatas(elemSz, cnt uint64) [][]byte {
	datas := make([][]byte, cnt)
	for i := range datas {
		data := make([]byte, elemSz)
		rand.Read(data)
		datas[i] = data
	}

	return datas
}

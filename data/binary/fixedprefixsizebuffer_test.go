package binary

import (
	. "github.com/kapitan-k/goutilities/test"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestFixedPrefixSizeBuffer(t *testing.T) {
	cnt := uint64(rand.Int63n(1024))
	elemSz := uint64(rand.Int63n(1024))

	if cnt < 4 {
		cnt = 4
	}

	if elemSz < 4 {
		elemSz = 4
	}

	datas := RandomDatas(elemSz, cnt)

	fpb := FixedPrefixSizeBufferCreate(elemSz, cnt)
	require.Equal(t, cnt, fpb.Cnt())

	fpb.SetAt(datas[0], 0)

	cdata := fpb.At(0)
	require.Equal(t, len(datas[0]), len(cdata))
	require.Equal(t, datas[0], cdata)

	for i, data := range datas {
		fpb.SetAt(data, uint64(i))
	}

	for i, data := range datas {
		require.Equal(t, data, fpb.At(uint64(i)))
	}

	fpb.SetAt(datas[3], uint64(3))

	for i, data := range datas {
		require.Equal(t, data, fpb.At(uint64(i)))
	}

	require.Equal(t, cnt, fpb.Cnt())

	lpre := len(fpb)
	cnt++
	fpb = fpb.Append(datas[4])
	require.Equal(t, lpre+int(elemSz), len(fpb))

	for i, data := range datas {
		require.Equal(t, data, fpb.At(uint64(i)))
	}
	require.Equal(t, cnt, fpb.Cnt())

	require.Equal(t, len(datas[4]), len(fpb.At(uint64(len(datas)))))
	require.Equal(t, datas[4], fpb.At(uint64(len(datas))))

	i := 0
	fnItr := func(v []byte) bool {
		require.Equal(t, v, fpb.At(uint64(i)))
		require.Equal(t, v, datas[i])
		i++
		if i == len(datas) {
			return false
		}
		return true
	}

	fpb.Iterate(fnItr)

	pdatas := fpb.Datas()
	require.Equal(t, datas, pdatas[:len(datas)])

}

func Test(t *testing.T) {
	cnt := uint64(rand.Int63n(1024))
	apb := ArrayLenPrefixedBufferCreate(0)
	datas := RandomDatasWithRandomElemSizePrefixed(cnt)
	for _, data := range datas {
		apb = apb.AppendArrayLenPrefixedData(data)
	}

	require.Equal(t, cnt, apb.Cnt())

	i := 0
	fnItr := func(v []byte) bool {
		require.Equal(t, v, datas[i][ArrayLenByteSz:])
		i++
		return true
	}

	apb.Iterate(fnItr)

	pdatas := apb.Datas()
	for i, data := range datas {
		require.Equal(t, data[ArrayLenByteSz:], pdatas[i])
	}

	apb = apb.Append(datas[0][ArrayLenByteSz:])
	datas = append(datas, datas[0])

	i = 0
	fnItr = func(v []byte) bool {
		require.Equal(t, v, datas[i][ArrayLenByteSz:])
		i++
		return true
	}

	apb.Iterate(fnItr)

	apb = apb.AppendMulti(datas)
	datas = append(datas, datas...)

	i = 0
	fnItr = func(v []byte) bool {
		if i < len(datas)/2 {
			require.Equal(t, v, datas[i][ArrayLenByteSz:])
		} else {
			require.Equal(t, v, datas[i])
		}

		i++
		return true
	}

	apb.Iterate(fnItr)

}

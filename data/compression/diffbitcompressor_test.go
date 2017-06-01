package math

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

// TODO fuzz
func TestDiff(t *testing.T) {
	{
		bs := make(Uint64Bitset, 16)
		bs.SetBit(64)
		require.True(t, bs.CheckIsBitSet(64))
		log.Println("bs", bs)
	}

	inOld := []uint64{1, 2, 3, 4, 5, 6, 7}
	inNew := []uint64{1, 2, 3, 4, 29, 8, 7}
	tgt := []uint64{0, 0}
	tgtUncomp := make([]uint64, len(inNew))

	byteSize := DiffBitUint64CompressorByteSize(inNew, inOld)
	require.Equal(t, 16+int(Uint64BitsetCntNeeded(uint64(len(inNew)))*8), int(byteSize))

	bs := make(Uint64Bitset, 8)

	cntDiffComp := DiffBitUint64CompressorCompressFromTo(tgt, bs, inNew, inOld)
	require.Equal(t, 2, int(cntDiffComp))

	log.Println(tgt)
	log.Println("bs", bs)

	cntDiffUncomp := DiffBitUint64CompressorDecompressTo(tgtUncomp, bs, tgt, inOld)
	log.Println("tgtUncomp", tgtUncomp)
	require.Equal(t, 2, int(cntDiffUncomp))
}

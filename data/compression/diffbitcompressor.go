package math

import (
	. "github.com/kapitan-k/goutilities"
	. "github.com/kapitan-k/goutilities/bitsnbytes"
	"unsafe"
)

//
// Used to compress very similar sets of uint64 data somewhat efficiently
//

func DiffBitUint64PlainDataDiffByteSize(inNew, inOld []uint64) (size uint64) {
	for i, n := range inNew {
		if n != inOld[i] {
			size += 8
		}
	}
	return
}

func DiffBitUint64CompressorByteSize(inNew, inOld []uint64) (size uint64) {
	l := uint64(len(inNew))
	return Uint64BitsetByteSizeNeeded(l) + DiffBitUint64PlainDataDiffByteSize(inNew, inOld)
}

func DiffBitUint64CompressorUncompressTo(tgt []uint64, bs Uint64Bitset, inSet, inOld []uint64) (cntDiff uint64) {
	for i, o := range inOld {
		if bs.CheckIsBitSet(uint64(i)) {
			tgt[i] = inSet[cntDiff]
			cntDiff++
		} else {
			tgt[i] = o
		}
	}

	return
}

func DiffBitUint64CompressorCompressFromTo(tgt []uint64, bs Uint64Bitset, inNew, inOld []uint64) (cntDiff uint64) {
	if len(inNew) < len(inOld) {
		inOld = inOld[:len(inNew)]
	}

	for i, n := range inNew {
		o := inOld[i]
		if n != o {
			tgt[cntDiff] = n
			bs.SetBit(uint64(i))
			cntDiff++
		}
	}

	return
}

type Uint64Bitset []uint64

func (self Uint64Bitset) SetBit(pos uint64) {
	Uint64BitsetSetBit(self, pos)
}

func (self Uint64Bitset) CheckIsBitSet(pos uint64) bool {
	return Uint64BitsetCheckIsBitSet(self, pos)
}

func Uint64BitsetSetBit(bs Uint64Bitset, pos uint64) {
	bs[pos/64] |= 1 << (pos % 64)
}

func Uint64BitsetCheckIsBitSet(bs Uint64Bitset, pos uint64) bool {
	d := pos % 64
	v := pos / 64
	return CHECK_IS_BIT(bs[v], d)
}

func Uint64BitsetCntNeeded(cntElems uint64) (cnt uint64) {
	d := cntElems % 64
	if d != 0 {
		cnt++
	}

	cnt += cntElems / 64
	return
}

func Uint64BitsetByteSizeNeeded(cntElems uint64) (size uint64) {
	return Uint64BitsetCntNeeded(cntElems) * 8
}

func DiffBitUint64CompressorInfoFromUintptr(ptr uintptr, cntElems uint64) (bs Uint64Bitset, data []uint64) {
	bscnt := Uint64BitsetCntNeeded(cntElems)
	bs = Uint64Bitset(UnsafeToUint64Slice(unsafe.Pointer(ptr), bscnt))
	data = UnsafeToUint64Slice(unsafe.Pointer(ptr+uintptr(bscnt*8)), cntElems)
	return
}

func DiffBitUint64CompressorInfoNew(cntElems uint64) (bs Uint64Bitset, data []uint64) {
	bscnt := Uint64BitsetCntNeeded(cntElems)
	data = make([]uint64, cntElems+bscnt)
	ptr := uintptr(unsafe.Pointer(&data[0]))
	bs = Uint64Bitset(UnsafeToUint64Slice(unsafe.Pointer(ptr), bscnt))
	data = UnsafeToUint64Slice(unsafe.Pointer(ptr+uintptr(bscnt*8)), cntElems)
	return
}

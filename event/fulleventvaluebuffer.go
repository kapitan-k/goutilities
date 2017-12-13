package event

import (
	. "github.com/kapitan-k/goutilities/data/binary"
)

type FullEventValueBuffer struct {
	FixedPrefixSizeBuffer
}

func FullEventValueBufferCreateFixed(cnt uint64) FullEventValueBuffer {
	return FullEventValueBuffer{
		FixedPrefixSizeBuffer: FixedPrefixSizeBufferCreate(FullEventHeaderByteSz, cnt),
	}
}

// valueSize is the size of the values
func FullEventValueBufferCreateFixedWithValueSize(valueSize, cnt uint64) FullEventValueBuffer {
	return FullEventValueBuffer{
		FixedPrefixSizeBuffer: FixedPrefixSizeBufferCreate(FullEventHeaderByteSz+valueSize, cnt),
	}
}

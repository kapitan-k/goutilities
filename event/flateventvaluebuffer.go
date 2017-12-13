package event

import (
	. "github.com/kapitan-k/goutilities/data/binary"
)

type FlatEventValueBuffer struct {
	FixedPrefixSizeBuffer
}

func FlatEventValueBufferCreate(cnt uint64) FlatEventValueBuffer {
	return FlatEventValueBuffer{
		FixedPrefixSizeBuffer: FixedPrefixSizeBufferCreate(ObjectHeaderByteSz, cnt),
	}
}

package iterator

type FnMoreData func(k []byte) (data []byte, err error)

type KVIteratorRequesterSeeker interface {
	SeekToFirst(k []byte) (data []byte, err error)
	SeekToLast(k []byte) (data []byte, err error)
	Seek(k []byte) (data []byte, err error)
	SeekForPrev(k []byte) (data []byte, err error)
}

type KVIteratorRequester interface {
	KVIteratorRequesterSeeker

	Next(kLast []byte) (data []byte, err error)
	Prev(kLast []byte) (data []byte, err error)
}

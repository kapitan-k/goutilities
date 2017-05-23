package external

type FnMoreData func(k []byte) (data []byte, err error)

type KVIteratorRequester interface {
	SeekToFirst(k []byte) (data []byte, err error)
	SeekToLast(k []byte) (data []byte, err error)
	SeekTo(k []byte) (data []byte, err error)
	SeekForPrev(k []byte) (data []byte, err error)

	Next(kLast []byte) (data []byte, err error)
	Prev(kLast []byte) (data []byte, err error)
}

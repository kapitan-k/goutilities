package iterator

// An iterator for key values, which is loosely based on the rocksdb iterator
type KVIterator interface {
	IteratorBase
	ValidTo(data []byte) bool
	SeekTo(data []byte)
	SeekForPrev(data []byte)
	KeyValue() (k, v []byte)
	Next()
	Prev()
}

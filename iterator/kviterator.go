package iterator

// An iterator for key values, which is loosely based on the rocksdb iterator
type KVIterator interface {
	IteratorBase
	Seek(k []byte)
	SeekForPrev(k []byte)
	KeyValue() (k, v []byte)
	Next()
	Prev()
}

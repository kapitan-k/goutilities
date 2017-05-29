package iterator

type KVBaseIterator interface {
	IteratorValidator
	Key() (k []byte)
	Value() (v []byte)
	KeyValue() (k, v []byte)
	Next()
}

// An iterator for key values, which is loosely based on the rocksdb iterator
type KVIterator interface {
	IteratorBase
	IteratorSeeker
	KVBaseIterator
	Prev()
}

type KBaseIterator interface {
	IteratorValidator
	Key() (k []byte)
	Next()
}

type VBaseIterator interface {
	IteratorValidator
	Value() (vk []byte)
	Next()
}

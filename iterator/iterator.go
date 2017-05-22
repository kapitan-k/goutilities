package iterator

type IteratorBase interface {
	Valid() bool
	SeekToFirst()
	SeekToLast()
	Close()
}

package iterator

type IteratorDirection uint64

const (
	IteratorDirection_Asc  = IteratorDirection(1)
	IteratorDirection_Desc = IteratorDirection(2)
)

type IteratorValidator interface {
	Valid() bool
}

type IteratorBase interface {
	Err() error
	Close()
}

type IteratorSeeker interface {
	SeekToFirst(k []byte)
	SeekToLast(k []byte)
	Seek(k []byte)
	SeekForPrev(k []byte)
}

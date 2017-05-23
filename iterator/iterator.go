package iterator

type IteratorDirection uint64

const (
	IteratorDirection_Asc  = IteratorDirection(1)
	IteratorDirection_Desc = IteratorDirection(2)
)

type IteratorBase interface {
	Err() error
	Valid() bool
	SeekToFirst(k []byte)
	SeekToLast(k []byte)
	Close()
}

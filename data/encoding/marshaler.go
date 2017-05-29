package encoding

type Marshaler interface {
	Size() int
	MarshalTo(data []byte) (int, error)
	Marshal() (data []byte, err error)
}

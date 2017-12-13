package encoding

// Marshaler is used to encode data to a byte slice.
type Marshaler interface {
	Size() int
	// MarshalTo serializes the Marshaler to data which
	// must be >= Size().
	MarshalTo(data []byte) (size int, err error)
	// MarshalTo serializes the Marshaler to data which
	// is newly allocated.
	Marshal() (data []byte, err error)
}

// Unmarshaler is used to decode data to a byte slice.
type Unmarshaler interface {
	// Unmarshal deserializes data to Unmarshaler.
	Unmarshal(data []byte) (err error)
}

// MarshalSerializer is a Marshaler and Unmarshaler.
type MarshalSerializer interface {
	Marshaler
	Unmarshaler
}

// ZeroMarshalSerializer s methods do nothing.
type ZeroMarshalSerializer struct{}

// Size returns 0.
func (self ZeroMarshalSerializer) Size() int {
	return 0
}

// MarshalTo returns 0, nil.
func (self ZeroMarshalSerializer) MarshalTo(data []byte) (size int, err error) {
	return
}

// MarshalTo returns nil, nil.
func (self ZeroMarshalSerializer) Marshal() (data []byte, err error) {
	return
}

// Unmarshal deserializes data to Unmarshaler.
func (self ZeroMarshalSerializer) Unmarshal(data []byte) (err error) {
	return
}

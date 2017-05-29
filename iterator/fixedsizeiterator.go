package iterator

type BasePositionedFormat struct {
	data []byte
	pos  uint64
}

type FixedSizeIterator struct {
	BasePositionedFormat
	kLen uint64
}

// can be used a s key or value iterator
func FixedSizeIteratorCreate(data []byte, kLen uint64) (self FixedSizeIterator) {
	self.data = data
	self.kLen = kLen
	return
}

func (self *FixedSizeIterator) Valid() bool {
	return self.pos < uint64(len(self.data))
}

func (self *FixedSizeIterator) Key() (k []byte) {
	return self.Value()
}

func (self *FixedSizeIterator) Value() (v []byte) {
	if self.pos == uint64(len(self.data)) {
		return nil
	}

	return self.data[self.pos : self.pos+self.kLen]
}

func (self *FixedSizeIterator) Next() {
	if self.pos == uint64(len(self.data)) {
		return
	}
	self.pos += self.kLen
}

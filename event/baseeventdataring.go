package event

func isPowerOfTwo(v uint64) bool {
	if v == 0 {
		return false
	}

	return (v & (v - 1)) == 0
}

// BaseEventDataRing is a ring of EventData
type BaseEventDataRing struct {
	datas    []BaseEventData
	writePos uint64
	readPos  uint64
}

func BaseEventDataRingCreate(cap int) (self BaseEventDataRing) {
	if !isPowerOfTwo(uint64(cap)) {
		panic("cap is not power of two")
	}
	self.datas = make([]BaseEventData, cap)
	return
}

// Offer puts data in the ring at current writePos if self.Free() > 0
// and increses the writePos.
func (self *BaseEventDataRing) Offer(data BaseEventData) (isIn bool) {
	if self.Free() == 0 {
		return false
	}
	self.datas[self.writePos%uint64(len(self.datas))] = data
	self.writePos++
	return true
}

// Pop returns the element at the current readPos if self.Empty() == false
// and increases the readPos.
func (self *BaseEventDataRing) Pop(ed *BaseEventData) (hasEvent bool) {
	if self.Empty() {
		return
	}
	rPos := self.readPos
	self.readPos++
	*ed = self.datas[rPos%uint64(len(self.datas))]
	return true
}

// Get returns the element at the current readPos if self.Empty() == false
func (self *BaseEventDataRing) Get(ed *BaseEventData) (hasEvent bool) {
	if self.Empty() {
		return
	}
	*ed = self.datas[self.readPos%uint64(len(self.datas))]
	return true
}

func (self *BaseEventDataRing) At() *BaseEventData {
	if self.Empty() {
		return nil
	}

	return &self.datas[self.readPos%uint64(len(self.datas))]
}

// IncReadPos pops cnt events. (adds cnt to self.readPos) panics if self.Cnt() < cnt
func (self *BaseEventDataRing) IncReadPos(cnt uint64) {
	if self.Cnt() < cnt {
		panic("self.Cnt() < cnt")
	}
	self.readPos += cnt
}

// Cnt returns the cnt of the elements in the ring.
func (self *BaseEventDataRing) Cnt() uint64 {
	return self.writePos - self.readPos
}

// Empty returns whether the ring is empty.
func (self *BaseEventDataRing) Empty() bool {
	return self.writePos == self.readPos
}

// Free returns the freespace in the ring.
func (self *BaseEventDataRing) Free() uint64 {
	return uint64(len(self.datas)) - self.Cnt()
}

// Reset sets writePos and readPos to zero and
// sets all element of its datas to nil.
func (self *BaseEventDataRing) Reset() {
	self.writePos = 0
	self.readPos = 0
	for i := range self.datas {
		self.datas[i] = BaseEventData{}
	}
}

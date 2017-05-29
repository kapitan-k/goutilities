package iterator

import (
	. "github.com/kapitan-k/goutilities/data"
)

type FnExternalKVRequestIteratorNext func(data []byte, pos int) (k, v []byte, posNew int)

type KVRequestIterator struct {
	requester  KVIteratorRequester
	fnNext     FnExternalKVRequestIteratorNext
	fnFreeData FnFreeData

	data []byte
	pos  int

	kLast []byte
	vLast []byte

	err error
}

func KVRequestIteratorCreate(requester KVIteratorRequester, fnFreeData FnFreeData) (self KVRequestIterator) {
	self.requester = requester
	self.fnFreeData = fnFreeData
	return
}

func (self *KVRequestIterator) Err() error {
	return self.err
}

func (self *KVRequestIterator) Valid() bool {
	return self.err == nil && self.pos < len(self.data)
}

func (self *KVRequestIterator) SeekToFirst(k []byte) {
	self.requestInternal(self.requester.SeekToFirst, k)
}

func (self *KVRequestIterator) SeekToLast(k []byte) {
	self.requestInternal(self.requester.SeekToLast, k)
}

func (self *KVRequestIterator) Seek(k []byte) {
	self.requestInternal(self.requester.Seek, k)
}

func (self *KVRequestIterator) SeekForPrev(k []byte) {
	self.requestInternal(self.requester.SeekForPrev, k)
}

func (self *KVRequestIterator) Key() (k []byte) {
	return self.kLast
}

func (self *KVRequestIterator) Value() (v []byte) {
	return self.vLast
}

func (self *KVRequestIterator) KeyValue() (k, v []byte) {
	return self.kLast, self.vLast
}

func (self *KVRequestIterator) Next() {
	if self.Valid() {
		self.kLast, self.vLast, self.pos = self.fnNext(self.data, self.pos)
		return
	} else if self.err != nil {
		sdata, err := self.requester.Next(self.kLast)
		if err != nil {
			self.err = err
			return
		}
		self.fnFreeData(self.data)
		self.data = sdata
	}

}

// with direction DESC
func (self *KVRequestIterator) Prev() {
	self.Next()
}

func (self *KVRequestIterator) requestInternal(fn FnMoreData, k []byte) {
	sdata, err := fn(k)
	if err != nil {
		self.err = err
		return
	}

	self.data = sdata
}

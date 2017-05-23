package external

import (
	//. "github.com/kapitan-k/goutilities/data"
	. "goutilities/data"
)

type FnExternalExternalKVIteratorNext func(data []byte, pos int) (k, v []byte, posNew int)

type ExternalKVIterator struct {
	requester  KVIteratorRequester
	fnNext     FnExternalExternalKVIteratorNext
	fnFreeData FnFreeData

	data []byte
	pos  int

	kLast []byte
	vLast []byte

	err error
}

func ExternalKVIteratorCreate(requester KVIteratorRequester, fnFreeData FnFreeData) (self ExternalKVIterator) {
	self.requester = requester
	self.fnFreeData = fnFreeData
	return
}

func (self *ExternalKVIterator) Error() error {
	return self.err
}

func (self *ExternalKVIterator) Valid() bool {
	return self.err == nil && self.pos < len(self.data)
}

func (self *ExternalKVIterator) SeekToFirst(k []byte) {
	self.requestInternal(self.requester.SeekToFirst, k)
}

func (self *ExternalKVIterator) SeekToLast(k []byte) {
	self.requestInternal(self.requester.SeekToLast, k)
}

func (self *ExternalKVIterator) SeekTo(k []byte) {
	self.requestInternal(self.requester.SeekTo, k)
}

func (self *ExternalKVIterator) SeekForPrev(k []byte) {
	self.requestInternal(self.requester.SeekForPrev, k)
}

func (self *ExternalKVIterator) Next() {
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
func (self *ExternalKVIterator) Prev() {
	self.Next()
}

func (self *ExternalKVIterator) requestInternal(fn FnMoreData, k []byte) {
	sdata, err := fn(k)
	if err != nil {
		self.err = err
		return
	}

	self.data = sdata
}

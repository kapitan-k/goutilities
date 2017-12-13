package event

import (
	"fmt"
	. "github.com/kapitan-k/goutilities/unsafe"
	"math"
	"sync/atomic"
	"time"
)

const (
	topicid_seq_bits             uint64 = 40
	topicid_context_type_bits    uint64 = 8
	topicid_context_subtype_bits uint64 = 16

	topicid_context_type_shift    uint64 = topicid_seq_bits + topicid_context_subtype_bits
	topicid_context_subtype_shift uint64 = topicid_seq_bits

	topicid_context_type_mask    uint64 = ^(^0 << topicid_context_type_bits) << topicid_context_type_shift
	topicid_context_subtype_mask uint64 = ^(^0 << topicid_context_subtype_bits) << topicid_seq_bits
	topicid_seq_mask             uint64 = ^(^0 << topicid_seq_bits) << 0
)

const TopicIDByteSz = 8
const TopicIDMax = TopicID(math.MaxInt64)

// TopicID is used to identify a unique topic. It contains a ContextType(), Time() in ms and a sequence Seq()
type TopicID uint64

func TopicIDCreate(contextType uint8, contextSubType uint16, seq uint64) (topicID TopicID) {
	topicID.SetContextType(uint64(contextType))
	topicID.SetContextSubType(uint64(contextSubType))
	topicID.SetSeq(seq)
	return
}

func (self *TopicID) String() string {
	return fmt.Sprint("TopicID", *self, " contextType ", self.ContextType(), " contextSubType ", self.ContextSubType(), " seq ", self.Seq())
}

func (self *TopicID) AsSlice() []byte {
	return ByteSliceFromUint64(uint64(*self))
}

func (self *TopicID) Next() {
	self.SetSeqNext()
}

func (self *TopicID) ContextType() uint64 {
	return (uint64(*self) & topicid_context_type_mask) >> topicid_context_type_shift
}

func (self *TopicID) ContextSubType() uint64 {
	return (uint64(*self) & topicid_context_subtype_mask) >> topicid_context_subtype_shift
}

func (self *TopicID) Seq() uint64 {
	return (uint64(*self) & topicid_seq_mask)
}

func (self *TopicID) SetSeqNext() {
	self.SetSeq(self.Seq() + 1)
}

func (self *TopicID) SetContextType(val uint64) {
	pp := uint64(*self)
	pp = pp & ^topicid_context_type_mask
	val = val << topicid_context_type_shift
	*self = TopicID(pp | val)

}

func (self *TopicID) SetContextSubType(val uint64) {
	pp := uint64(*self)
	pp = pp & ^topicid_context_subtype_mask
	val = val << topicid_context_subtype_shift
	*self = TopicID(pp | val)
}

func (self *TopicID) SetSeq(val uint64) {
	pp := uint64(*self)
	pp = pp & ^topicid_seq_mask
	val = val & topicid_seq_mask
	*self = TopicID(pp | val)
}

func (self *TopicID) Atomic() TopicID {
	return TopicID(atomic.LoadUint64((*uint64)(self)))
}

func (self *TopicID) AtomicSet(val TopicID) {
	atomic.StoreUint64((*uint64)(self), uint64(val))
}

func (self *TopicID) AtomicSetSeq(val uint64) {
	ppcur := TopicID(*self)
	ppcur.SetSeq(val)
	atomic.StoreUint64((*uint64)(self), uint64(ppcur))
}

func (self *TopicID) AtomicNext() (cur, next TopicID) {
	for {
		cur = self.Atomic()
		next = cur
		next.SetSeqNext()
		if atomic.CompareAndSwapUint64((*uint64)(self), uint64(cur), uint64(next)) {
			return
		}
		time.Sleep(time.Nanosecond * 64)
	}
}

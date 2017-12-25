package event

import (
	"fmt"
	"math"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	topiceventid_time_bits uint64 = 52
	topiceventid_seq_bits  uint64 = 12

	topiceventid_seq_max uint64 = (^(^0 << topiceventid_seq_bits) << 0) - 1

	topiceventid_time_shift uint64 = topiceventid_seq_bits
	topiceventid_seq_shift  uint64 = topiceventid_seq_bits

	topiceventid_time_mask uint64 = ^(^0 << topiceventid_time_bits) << topiceventid_seq_bits
	topiceventid_seq_mask  uint64 = ^(^0 << topiceventid_seq_bits) << 0

	topiceventid_custom_tim_beg uint64 = 0
)

const EventIDByteSz = 8
const EventIDMax = EventID(math.MaxUint64)

// EventID is used to identify a unique event of a topic.
type EventID uint64

func (self *EventID) String() string {
	return fmt.Sprint("EventID", *self, " time ", self.TimeUsec(), " seq ", self.Seq(), time.Unix(int64(self.TimeUsec()/1000000), 0).UTC())
}

func (self *EventID) Next(timeMicros uint64) {
	timLast := self.TimeUsec()
	seqLast := self.Seq()

L:
	if timeMicros > timLast {
		seqLast = 0
	} else if timeMicros == timLast {
		seqLast++
		if seqLast == topiceventid_seq_max {
			time.Sleep(time.Microsecond)
			goto L
		}
	}

	self.SetTimeUsec(timeMicros)
	self.SetSeq(seqLast)

	return
}

func (self EventID) TimeUsec() uint64 {
	return ((uint64(self) & topiceventid_time_mask) >> topiceventid_time_shift) + topiceventid_custom_tim_beg
}

func (self *EventID) Seq() uint64 {
	return (uint64(*self) & topiceventid_seq_mask)
}

func (self *EventID) SetSeqNext() {
	self.SetSeq(self.Seq() + 1)
}

func (self *EventID) SetTimeUsec(val uint64) {
	val -= topiceventid_custom_tim_beg
	pp := uint64(*self)
	pp = pp & ^topiceventid_time_mask
	val = val << topiceventid_time_shift
	*self = EventID(pp | val)
}

func (self *EventID) SetTimeUsecMax() {
	*self = EventID(uint64(*self) | topiceventid_time_mask)
}

func (self *EventID) SetSeq(val uint64) {
	pp := uint64(*self)
	pp = pp & ^topiceventid_seq_mask
	val = val & topiceventid_seq_mask
	*self = EventID(pp | val)
}

func (self *EventID) AddTimeUsec(val uint64) {
	self.SetTimeUsec(self.TimeUsec() + val)
}

func (self *EventID) Atomic(tid *uint64) uint64 {
	return atomic.LoadUint64(tid)
}

func (self *EventID) AtomicSet(val uint64) {
	atomic.StoreUint64((*uint64)(self), val)
}

func (self *EventID) AtomicSetTimeUsec(val uint64) {
	val -= topiceventid_custom_tim_beg
	teid := EventID(*self)
	teid.SetTimeUsec(val)
	atomic.StoreUint64((*uint64)(self), uint64(teid))
}

func (self *EventID) AtomicAddTimeUsec(val uint64) {
	teid := EventID(*self)
	teid.SetTimeUsec(teid.TimeUsec() + val)
	atomic.StoreUint64((*uint64)(self), uint64(teid))
}

func (self *EventID) AtomicSetSeq(val uint64) {
	teid := EventID(*self)
	teid.SetSeq(val)
	atomic.StoreUint64((*uint64)(self), uint64(teid))
}

// FromSlice copies the data from data to EventID.
func (self *EventID) FromSlice(data []byte) {
	*self = *(*EventID)(unsafe.Pointer(&data[0]))
}

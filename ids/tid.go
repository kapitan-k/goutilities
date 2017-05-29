package ids

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"sync/atomic"
	"time"
)

// 12 seq bits ist max sequence generation of 4096-1 unique per micro if broken down to time unit
const (
	tid_time_bits     uint64 = 52
	tid_totalseq_bits uint64 = 12
	seq_max           uint64 = 16

	tid_time_shift uint64 = tid_totalseq_bits
	tid_seq_shift  uint64 = 8

	tid_time_mask          uint64 = ^(^0 << 52) << tid_totalseq_bits
	tid_seq_mask           uint64 = ^(^0 << 4) << 8
	tid_topic_seq_mask     uint64 = ^(^0 << 8) << 0
	tid_seq_max_mask       uint64 = ^(^0 << 12) << 0
	tid_topic_seq_max_mask uint64 = ^(^0 << 4) << 0

	Tid_time_mask uint64 = tid_time_mask
	Tid_seq_mask  uint64 = tid_seq_mask

	tid_custom_tim_beg uint64 = 0
)

func TidPrintln(tid uint64) {
	log.Println("tid ", " tim ", TidTim(tid), " seq ", TidSeq(tid))
}

func TidStr(tid uint64) string {
	return fmt.Sprint("tid", tid, " tim ", TidTim(tid), " seq ", TidSeq(tid), " topic seq ", TidTopicSeq(tid), time.Unix(int64(TidTim(tid)/1000000), 0).UTC())
}

func TidTim(pp uint64) uint64 {
	return ((pp & tid_time_mask) >> tid_time_shift) + tid_custom_tim_beg
}

func TidNext(tim, tidLast, seqTopic uint64) (tid uint64) {
	timLast := TidTim(tidLast)
	seqLast := TidSeq(tidLast)
L:
	if tim > timLast {
		seqLast = 1
	} else if tim == timLast {
		seqLast++
		if seqLast == seq_max {
			time.Sleep(time.Microsecond)
			seqLast = 1
			goto L
		}
	}

	TidSetTim(&tid, tim)
	TidSetSeq(&tid, seqLast)
	TidSetTopicSeq(&tid, seqTopic+1)

	return
}

func TidSeq(pp uint64) uint64 {
	return (pp & tid_seq_mask) >> tid_seq_shift
}

func TidTopicSeq(pp uint64) uint64 {
	return (pp & tid_topic_seq_mask)
}

func TidSetSeqNext(pp uint64) uint64 {
	ppx := pp
	TidSetSeq(&ppx, TidSeq(pp)+1)
	return ppx
}

func TidSeqIsNextOK(ppLast, ppNow uint64) bool {
	seqExpected := TidSeq(TidSetSeqNext(ppLast))
	seqIs := TidSeq(ppNow)
	if seqExpected == seqIs {
		return true
	}

	return false
}

func TidSetTim(tid *uint64, val uint64) {
	val -= tid_custom_tim_beg
	pp := uint64(*tid)
	pp = pp & ^tid_time_mask
	val = val << tid_time_shift
	*tid = (pp | val)
}

func TidSetTimMax(tid *uint64) {
	*tid = (*tid | tid_time_mask)
}

func TidSetSeq(tid *uint64, val uint64) {
	pp := uint64(*tid)
	pp = pp & ^tid_seq_mask
	val = val << tid_seq_shift
	*tid = (pp | val)
}

func TidSetTopicSeq(tid *uint64, val uint64) {
	pp := uint64(*tid)
	pp = pp & ^tid_topic_seq_mask
	val = val & tid_topic_seq_mask
	*tid = (pp | val)
}

func TidAddTim(tid *uint64, val uint64) {
	ppcur := *tid
	TidSetTim(&ppcur, TidTim(ppcur)+val)
	*tid = ppcur
}

func TidAtomic(tid *uint64) uint64 {
	return atomic.LoadUint64(tid)
}

func TidAtomicSet(tid *uint64, val uint64) {
	atomic.StoreUint64(tid, val)
}

func TidAtomicSetTim(tid *uint64, val uint64) {
	val -= tid_custom_tim_beg
	ppcur := *tid
	TidSetTim(&ppcur, val)
	atomic.StoreUint64(tid, ppcur)
}

func TidAtomicAddTim(tid *uint64, val uint64) {
	ppcur := *tid
	TidSetTim(&ppcur, TidTim(ppcur)+val)
	atomic.StoreUint64(tid, ppcur)
}

func TidAtomicSetSeq(tid *uint64, val uint64) {
	ppcur := *tid
	TidSetSeq(&ppcur, val)
	atomic.StoreUint64(tid, ppcur)
}

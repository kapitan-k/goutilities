package ids

import (
	"testing"
)

func TestTid(t *testing.T) {
	var tid, tid2 uint64
	timSet := uint64(1488375171096000)
	seqSet := uint64(14)
	topicSeqSet := uint64(255)

	TidSetTime(&tid, timSet)
	timGet := TidTime(tid)
	if timGet != timSet {
		t.Log("timGet != timSet", timGet, timSet)
		t.FailNow()
	}

	timGet = TidTime(tid)
	if timGet != timSet {
		t.Log("timGet != timSet", timGet, timSet, TidStr(tid))
		t.FailNow()
	}

	TidSetSeq(&tid, seqSet)
	TidSetTopicSeq(&tid, topicSeqSet)

	timGet = TidTime(tid)
	if timGet != timSet {
		t.Log("timGet != timSet", timGet, timSet, TidStr(tid))
		t.FailNow()
	}

	seqGet := TidSeq(tid)
	if seqGet != seqSet {
		t.Log("seqGet != seqSet", seqGet, seqSet, TidStr(tid))
		t.FailNow()
	}

	topicSeqGet := TidTopicSeq(tid)
	if topicSeqGet != topicSeqSet {
		t.Log("topicSeqGet != topicSeqSet", topicSeqGet, topicSeqSet, TidStr(tid))
		t.FailNow()
	}

	TidAddTime(&tid, 55)

	timGet = TidTime(tid)
	if timGet != timSet+55 {
		t.Log("timGet != timSet", timGet, timSet, TidStr(tid))
		t.FailNow()
	}

	seqGet = TidSeq(tid)
	if seqGet != seqSet {
		t.Log("seqGet != seqSet", seqGet, seqSet, TidStr(tid))
		t.FailNow()
	}

	TidSetSeq(&tid, 16)

	seqGet = TidSeq(tid)
	if seqGet != 0 {
		t.Log("seqGet != 0", seqGet, TidStr(tid))
		t.FailNow()
	}

	TidSetTopicSeq(&tid, topicSeqSet+1)
	topicSeqGet = TidTopicSeq(tid)
	if topicSeqGet != 0 {
		t.Log("topicSeqGet != 0", topicSeqGet, topicSeqSet, TidStr(tid))
		t.FailNow()
	}

	TidSetSeq(&tid2, 1)
	isSeqNextOk := TidSeqIsNextOK(tid, tid2)
	if !isSeqNextOk {
		t.Log("!isSeqNextOk", seqGet, TidStr(tid))
		t.FailNow()
	}

	TidAtomicSetTime(&tid, 987654)
	tid = TidAtomic(&tid)

	timGet = TidTime(tid)
	if timGet != 987654 {
		t.Log("timGet != timSet", timGet, timSet, TidStr(tid))
		t.FailNow()
	}
}

package iterator

const (
	DiffTypeNotExistInExpect = 1
	DiffTypeNotExistInActual = 2
	DiffTypeDifferKey        = 3
	DiffTypeDifferValue      = 4
)

// SingleKVDiff is a Diff for a single Key Value Tuple

// SingleKVTableDiffs = All the diffs for a single table/set/column family

// Diff = All the diffs between the DB in Path and the expected DB

type DiffIteratorProvider interface {
	DiffIteratorsAndNames() (iters []KVIterator, names []string)
}

func DiffDBs(dpExpect DiffIteratorProvider, dpActuals []DiffIteratorProvider, singleKVDiffsLimit uint64) (diffs []Diff, err error) {
	if len(dpActuals) == 0 {
		panic("len(dpActuals) == 0")
	}

	itersExpect, namesExpect := dpExpect.DiffIteratorsAndNames()
	l := len(itersExpect)
	if l != len(namesExpect) {
		panic("len(itersExpect) != len(namesExpect)")
	}

	for _, dpActual := range dpActuals {
		itersActual, namesActual := dpActual.DiffIteratorsAndNames()
		if len(itersActual) != len(namesActual) {
			panic("len(itersActual) != len(namesActual)")
		}
		if len(itersActual) != l {
			panic("len(itersActual) != len(itersExpect)")
		}
	}

	diffs = make([]Diff, len(dpActuals))
	for i, dpActual := range dpActuals {
		var tDiffs []SingleKVTableDiffs
		itersActual, namesActual := dpActual.DiffIteratorsAndNames()
		tDiffs, err = DiffKVIteratorsTotal(itersExpect, itersActual, namesActual, singleKVDiffsLimit)
		if err != nil {
			return
		}

		diffs[i].TableDiffs = tDiffs
	}

	return
}

func DiffKVIteratorsTotal(itersExpect, itersActual []KVIterator, tableNames []string, limit uint64) (tDiffs []SingleKVTableDiffs, err error) {
	tDiffs = make([]SingleKVTableDiffs, len(itersActual))
	for i, iterExpect := range itersExpect {
		var kvDiffs []SingleKVDiff
		iterActual := itersActual[i]
		kvDiffs, err = DiffKVIterators(iterExpect, iterActual, limit)
		if err != nil {
			return
		}

		tDiffs[i] = SingleKVTableDiffs{
			TableName: tableNames[i],
			KVDiffs:   kvDiffs,
		}
	}

	return
}

func DiffKVIterators(iterExpect, iterActual KVIterator, limit uint64) (diffs []SingleKVDiff, err error) {
	var cnt uint64
	var kExpect, vExpect []byte
	var kActual, vActual []byte

	iterActual.SeekToFirst(nil)
	isValidActual := iterActual.Valid()

	for iterExpect.SeekToFirst(nil); iterExpect.Valid() && cnt < limit; iterExpect.Next() {
		kExpect, vExpect = iterExpect.KeyValue()
		skExpect := string(kExpect)
		if !isValidActual {
			sd := SingleKVDiff{
				DiffType:      DiffTypeNotExistInExpect,
				Idx:           cnt,
				KeyExpected:   copyBuf(kExpect),
				ValueExpected: copyBuf(vExpect),
			}
			diffs = append(diffs, sd)
			cnt++
			continue
		}

		kActual, vActual = iterActual.KeyValue()

		skActual := string(kActual)
		if skExpect == skActual {
			svExpect := string(vExpect)
			svActual := string(vActual)
			if svExpect != svActual {
				sd := SingleKVDiff{
					DiffType:      DiffTypeDifferValue,
					Idx:           cnt,
					ValueActual:   copyBuf(vActual),
					ValueExpected: copyBuf(vExpect),
				}
				diffs = append(diffs, sd)
			}
			iterActual.Next()
			isValidActual = iterActual.Valid()

		} else if skExpect < skActual {
			sd := SingleKVDiff{
				DiffType:    DiffTypeNotExistInActual,
				Idx:         cnt,
				KeyActual:   copyBuf(kActual),
				KeyExpected: copyBuf(kExpect),
			}
			diffs = append(diffs, sd)

		} else if skExpect > skActual {
			sd := SingleKVDiff{
				DiffType:    DiffTypeNotExistInExpect,
				Idx:         cnt,
				KeyActual:   copyBuf(kActual),
				KeyExpected: copyBuf(kExpect),
			}
			diffs = append(diffs, sd)

			iterActual.Next()
			isValidActual = iterActual.Valid()
		}

		cnt++
	}

	for ; iterActual.Valid() && cnt < limit; iterActual.Next() {
		kActual, vActual = iterActual.KeyValue()
		sd := SingleKVDiff{
			DiffType:    DiffTypeNotExistInActual,
			Idx:         cnt,
			KeyActual:   copyBuf(kActual),
			ValueActual: copyBuf(vActual),
		}
		diffs = append(diffs, sd)
		cnt++
	}
	return
}

func kvToStrings(k, v []byte) (ks, vs string) {
	return string(k), string(v)
}

func copyBuf(in []byte) (out []byte) {
	out = make([]byte, len(in))
	copy(out, in)
	return
}

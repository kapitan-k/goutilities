package iterator

const (
	DiffTypeNotExistInExpect = 1
	DiffTypeNotExistInActual = 2
	DiffTypeDifferKey        = 3
	DiffTypeDifferValue      = 4
)

// Diff for a single Key Value Tuple
type SingleKVDiff struct {
	DiffType      int
	Idx           uint64
	KeyExpected   []byte
	ValueExpected []byte

	KeyActual   []byte
	ValueActual []byte
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

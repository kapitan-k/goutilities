package binary

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestEventBuffer(t *testing.T) {
	var kveb, kvebGet KVEventBufferHolder

	cnt := uint64(rand.Int63n(512))
	keySz := uint64(rand.Int63n(512))

	if cnt == 0 {
		cnt = 1
	}

	if keySz == 0 {
		keySz = 1
	}

	kveb.Keb.FixedKeySize = uint16(keySz)

	kveb.OffsetData = RandomDatas(keySz, 1)[0]
	kveb.Keys = RandomDatas(keySz, cnt)
	kveb.Values = RandomDatasWithRandomElemSizePrefixedMinSize(KVEventBufferHeaderByteSz+ArrayLenByteSz, cnt)

	evb := KVEventBuffer(KVEventBufferHolderToEventBuffer(&kveb, 0))

	KVEventBufferHolderFromEventBuffer(&kvebGet, evb)

	for i, key := range kveb.Keys {
		require.Equal(t, key, kvebGet.Keys[i])
	}

	for i, value := range kveb.Values {
		require.Equal(t, value, kvebGet.Values[i])
	}

}

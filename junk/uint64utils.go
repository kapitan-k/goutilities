package junk

import (
	"strconv"
	"strings"
)

// UInt64sToCommaSeparatedString returns a comma separated string
// comprising all vals.
func UInt64sToCommaSeparatedString(vals []uint64) string {
	b := make([]string, len(vals))
	for i, v := range vals {
		b[i] = strconv.FormatUint(v, 10)
	}

	return strings.Join(b, ",")
}

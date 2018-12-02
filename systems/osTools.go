package systems

import (
	"runtime"
	"strconv"
	"strings"
)

func GetMinimumVersion(v string) (uint64, error) {
	if v == "" {
		v = runtime.Version()
	}
	first := strings.IndexByte(v, '.')
	last := strings.LastIndexByte(v, '.')
	if first == last {
		return strconv.ParseUint(v[first+1:], 10, 64)
	}
	return strconv.ParseUint(v[first+1:last], 10, 64)
}

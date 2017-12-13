package junk

import (
	"errors"
	"strconv"
	"strings"
)

type BasicHost struct {
	Name string
	Port uint16
}

func BasicHostToString(host *BasicHost) string {
	return host.Name + ":" + strconv.FormatUint(uint64(host.Port), 10)
}

func BasicHostFromString(str string) (host BasicHost, err error) {
	idx := strings.LastIndex(str, ":")
	if idx == -1 {
		return host, errors.New("format")
	}
	host.Name = str[:idx]

	strPort := str[idx+1:]
	val, err := strconv.ParseInt(strPort, 10, 64)
	if err != nil {
		return host, err
	}
	host.Port = uint16(val)
	return host, nil
}

func BasicHostsToStrings(hosts []BasicHost) []string {
	strs := make([]string, len(hosts))
	for i, host := range hosts {
		strs[i] = BasicHostToString(&host)
	}
	return strs
}

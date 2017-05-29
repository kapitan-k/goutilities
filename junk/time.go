package junk

import (
	"time"
)

const DayMS = 86400000

const (
	TIME64_YEAR_BASE = 1900
)

// to enable easy querying of day, wday, month, year, hour, min, sec, TZ
// ignores msec, so it is not 100% exact and should be used "additionally"
type Time64 struct {
	Year  uint8
	Month uint8
	MDay  uint8
	WDay  uint8

	Hour uint8
	Min  uint8
	Sec  uint8
	TZ   uint8
}

func (self *Time64) FromUTC(timUtc int64, tz uint8) {
	now := time.Unix(int64(timUtc), 0)
	now = now.UTC()
	self.Year = uint8(now.Year() - TIME64_YEAR_BASE)
	self.Month = uint8(now.Month())
	self.MDay = uint8(now.Day())
	self.WDay = uint8(now.Weekday())
	self.Hour = uint8(now.Hour())
	self.Min = uint8(now.Minute())
	self.Sec = uint8(now.Second())
}

func TimeTodayUTCMs() int64 {
	now := time.Now().UTC()
	tim := now.Unix()
	secsToday := now.Second() + now.Minute()*60 + now.Hour()*3600 + (now.Nanosecond() / 1000000000)
	tim -= (int64)(secsToday)
	tim *= 1000
	return (int64)(tim)
}

func TimeNowUTCMicros() int64 {
	return (int64)(time.Now().UTC().UnixNano() / 1000)
}

func TimeNowUTCMs() int64 {
	return (int64)(time.Now().UTC().UnixNano() / 1000000)
}

func TimeNowUTCNano() int64 {
	return (int64)(time.Now().UTC().UnixNano())
}

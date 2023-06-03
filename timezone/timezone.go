package timezone

import (
	"time"
)

type Timezone struct {
	UtcOffset    string `json:"utc_offset"`
	Abbreviation string `json:"abbr"`
	IsDst        bool   `json:"is_dst"`
}

func GetTimezone() *Timezone {
	tz := &Timezone{}
	t := time.Now()
	utcOffset := t.Format("-07:00")
	tz.UtcOffset = utcOffset
	zone, _ := t.Zone()
	tz.Abbreviation = zone
	tz.IsDst = t.IsDST()

	return tz
}

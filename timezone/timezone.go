package timezone

import (
	"fmt"
	"runtime"
	"time"
)

type Timezone struct {
	UtcOffset    string `json:"utc_offset"`
	Abbreviation string `json:"abbr"`
	IsDst        bool   `json:"is_dst"`
	IanaName     string `json:"iana_name"`
}

func GetTimezone() *Timezone {
	tz := &Timezone{}
	t := time.Now()
	utcOffset := t.Format("-07:00")
	tz.UtcOffset = utcOffset
	zone, _ := t.Zone()
	tz.Abbreviation = zone
	tz.IsDst = t.IsDST()

	target := runtime.GOOS
	switch target {
	case "linux", "darwin":
		name, err := GetIanaNameForLinux()
		if err != nil {
			fmt.Println("Error reading IANA name from localtime", err)
			tz.IanaName = "Unknown/Unknown"
		} else {
			tz.IanaName = name
		}
	case "windows":
		name, err := GetIanaNameForWindows()
		if err != nil {
			fmt.Println("Error reading IANA name for Windows", err)
			tz.IanaName = "Unknown/Unknown"
		} else {
			tz.IanaName = name
		}
	default:
		fmt.Printf("Unsuported Platform '%v'\n", target)
		tz.IanaName = "Unknown/Unknown"
	}

	return tz
}

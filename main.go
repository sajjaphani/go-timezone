package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/sajjaphani/go-timezone/timezone"
)

func main() {
	tz := timezone.GetTimezone()
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(&tz); err != nil {
		log.Println(err)
	}
}

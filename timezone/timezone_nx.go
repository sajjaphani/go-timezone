package timezone

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func GetNameFromLocaltime() (string, error) {
	linkPath := "/etc/localtime"
	targetPath, err := os.Readlink(linkPath)
	if err != nil {
		return "", err
	}

	tzParts := strings.Split(targetPath, "/")
	if len(tzParts) < 3 {
		return "", errors.New("invalid timezone format")
	}

	continent, country := tzParts[len(tzParts)-2], tzParts[len(tzParts)-1]
	timezone := fmt.Sprintf("%s/%s", continent, country)

	// Load the location using the timezone value
	// Ensure valid IANA name
	_, err = time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}

	return timezone, nil
}

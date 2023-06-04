package timezone

import (
	"bufio"
	"embed"
	"encoding/json"
	"os/exec"
	"strings"
)

//go:embed resources/*.json
var Resources embed.FS

type WindowsZoneMapping struct {
	WindowsTimezone string
	IANATimezone    string
}

func GetIanaNameForWindowName(windowsName string) (string, error) {
	file, err := Resources.Open("resources/windows2iana.json")
	if err != nil {
		return "", err
	}

	r := bufio.NewReader(file)
	decoder := json.NewDecoder(r)
	var m map[string]string
	err = decoder.Decode(&m)

	return m[windowsName], err
}

func GetIanaNameForWindows() (string, error) {
	cmd := exec.Command("systeminfo")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	timezone := parseTimezone(string(output))

	return GetIanaNameForWindowName(timezone)
}

func parseTimezone(output string) string {
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Time Zone:") {
			// Extract the timezone value after the colon
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}

	return ""
}

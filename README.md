# Timezone

A [Go](https://golang.org/) project for finding timezone information. It retrieves the following information from the system.

- Offset from UTC: The offset represents the difference in hours and minutes from Coordinated Universal Time (UTC).
- Timezone Abbreviations: Timezones are often associated with three-letter abbreviations, such as "IST" (Indian Standard Time) or "EST" (Eastern Standard Time).
- Daylight Saving Time (DST): Many timezones observe Daylight Saving Time, which is the practice of adjusting the clocks forward by one hour during a specific period of the year.
- Timezone Name: A timezone is identified by a name that represents its location, such as "Asia/Kolkata" or "America/New_York"."

## Usage

Run the following command to see the JSON output

```sh
go run main.go
```

JSON output

```json
{
	"utc_offset": "+05:30",
	"abbr": "IST",
	"is_dst": false,
	"iana_name": "Asia/Kolkata"
}
```

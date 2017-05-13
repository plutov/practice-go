package webscraping

import "testing"
import "time"

import "strings"

var timezones = map[string]string{
	"UTC":  "UTC",
	"PDT":  "US/Pacific",
	"MDT":  "US/Mountain",
	"CDT":  "US/Central",
	"EDT":  "US/Eastern",
	"AKDT": "US/Alaska",
	"HAST": "US/Hawaii",
}

var tests = []string{
	"",
	"DT",
	"Alaska",
	"UTC",
	"HAST",
	"Apr. 19, 12:59:44 UTC",
	"US Naval Observatory",
	"AM",
}

func generateTimesForTimezone(timezone string, span int) ([]string, error) {
	result := []string{}

	tzName, ok := timezones[timezone]
	if !ok {
		return result, nil
	}

	loc, err := time.LoadLocation(tzName)
	if err != nil {
		return result, err
	}

	baseTime := time.Now().In(loc)

	for offset := -span; offset <= span; offset++ {
		offsetTime := baseTime.Add(time.Duration(offset) * time.Second)

		var format string
		if timezone == "UTC" {
			format = "Jan. 02, 15:04:05"
		} else {
			format = "Jan. 02, 03:04:05 PM"
		}

		formattedOffsetTime := offsetTime.Format(format) + " " + timezone
		result = append(result, formattedOffsetTime)
	}

	return result, nil
}

func TestGetTime(t *testing.T) {
	span := 10

	for _, tz := range tests {
		actual := GetTime(tz)

		expectedTimes, err := generateTimesForTimezone(tz, span)
		if err != nil {
			t.Fatalf("%s", err)
			continue
		}

		if len(expectedTimes) == 0 {
			expectedTimes = append(expectedTimes, "")
		}

		found := false
		for _, expected := range expectedTimes {
			if expected == actual {
				found = true
				break
			}
		}

		if !found {
			expectedList := strings.Join(expectedTimes, ", ")
			t.Fatalf("GetTime(\"%s\") expected one of %s, got %s", tz, expectedList, actual)
		}
	}
}

func BenchmarkGetTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tz := range tests {
			GetTime(tz)
		}
	}
}

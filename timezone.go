package goonung

import (
	"errors"
	"time"
)

// ConvertCurrentTimeToDifferentTimezone : convert current time to other timezone
func ConvertCurrentTimeToDifferentTimezone(now time.Time, location string) (time.Time, string, error) {
	timezone, err := time.LoadLocation(location)

	if err != nil {
		return now, location, err
	}

	now = now.In(timezone)

	return now, location, nil
}

// GetCurrentTime : get current time
func GetCurrentTime() (time.Time, error) {
	now := time.Now()

	return now, nil
}

// GetCurrentTimezone : get current timezone
func GetCurrentTimezone(now time.Time) (string, error) {
	zone, _ := now.Zone()

	return zone, nil
}

// GetCurrentTimeInVariousTimezone : get current time on various timezone
func GetCurrentTimeInVariousTimezone(now time.Time, locations []string) (map[string]time.Time, error) {
	timezones := make(map[string]time.Time)

	if len(locations) == 0 {
		return timezones, errors.New("empty or invalid timezone")
	}

	for _, location := range locations {
		differNow, location, err := ConvertCurrentTimeToDifferentTimezone(now, location)

		if err != nil {
			return timezones, err
		}

		timezones[location] = differNow
	}

	return timezones, nil
}

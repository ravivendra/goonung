package goonung

import (
	"testing"
	"time"
)

func TestConvertCurrentTimeToDifferentTimezone(t *testing.T) {
	if _, _, err := ConvertCurrentTimeToDifferentTimezone(time.Now(), "Europe/Jakarta"); err.Error() == "Unknown time zone Europe/Jakarta" {
		t.Errorf("Invalid Current Time to different Timezone : %s", err.Error())
	}

	if _, _, err := ConvertCurrentTimeToDifferentTimezone(time.Now(), "Europe/Berlin"); err != nil {
		t.Errorf("Get Current Time to different Timezone : %s", err.Error())
	}
}

func TestGetCurrentTime(t *testing.T) {
	if _, err := GetCurrentTime(); err != nil {
		t.Errorf("Get Current Time : %s", err.Error())
	}
}

func TestGetCurrentTimezone(t *testing.T) {
	if _, err := GetCurrentTimezone(time.Now()); err != nil {
		t.Errorf("Get Current Timezone : %s", err.Error())
	}
}

func TestGetCurrentTimeInVariousTimezone(t *testing.T) {
	locations := []string{}

	if _, err := GetCurrentTimeInVariousTimezone(time.Now(), locations); err.Error() == "Empty or invalid timezone" {
		t.Errorf("Empty Current Time in various timezone : %s", err.Error())
	}

	locations = []string{"Asia/Berlin"}

	if _, err := GetCurrentTimeInVariousTimezone(time.Now(), locations); err.Error() == "Unknown time zone Asia/Berlin" {
		t.Errorf("Invalid Current Time in various timezone : %s", err.Error())
	}

	locations = []string{"Asia/Tokyo", "Europe/Berlin"}

	if _, err := GetCurrentTimeInVariousTimezone(time.Now(), locations); err != nil {
		t.Errorf("Get Current Time in various timezone : %s", err.Error())
	}
}

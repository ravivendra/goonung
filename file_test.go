package goonung

import "testing"

func TestRead(t *testing.T) {
	if _, err := Read("sample/logs.txt", 0400); err.Error() == "No such file or directory" {
		t.Errorf("Invalid file : %s", err.Error())
	}

	if _, err := Read("sample/log.txt", 0400); err == nil {
		// t.Errorf("Valid file")
	}
}

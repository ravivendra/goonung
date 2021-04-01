package goonung

import (
	"testing"
)

func TestRotation(t *testing.T) {
	if err := FileRotation("false-sample", "txt", "tar.gz"); err.Error() == "no such file or directory" {
		t.Errorf("Invalid required directory : %s", err.Error())
	}

	if err := FileRotation("sample", "txt", "tar.gz"); err != nil {
		t.Errorf("Rotate file using TAR.GZ : %s", err.Error())
	}

	if err := FileRotation("sample", "go", "tar.gz"); err.Error() == "Empty files for your processing stage" {
		t.Errorf("Empty files with this requirement : %s", err.Error())
	}
}

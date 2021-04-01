package rotation

import "testing"

func TestDelete(t *testing.T) {
	if err := Delete("log.log"); err.Error() == "no such file or directory" {
		t.Errorf("Delete invalid file : %s", err.Error())
	}
}

func TestRotation(t *testing.T) {
	var files []Stat

	if err := Rotation("../sample", files, "tar.gz"); err.Error() == "Empty files for your processing stage" {
		t.Errorf("Empty files with this requirement : %s", err.Error())
	}

	files = append(files, Stat{FilePath: "../sample/log.txt", FileName: "log.txt"})

	if err := Rotation("../sample", files, "tar.gz"); err != nil {
		t.Errorf("Rotate file using TAR.GZ : %s", err.Error())
	}

	if err := Rotation("../sample", files, "zip"); err != nil {
		t.Errorf("Rotate file using ZIP : %s", err.Error())
	}
}

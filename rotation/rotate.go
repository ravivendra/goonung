package rotation

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ravivendra/goonung/tarball"
)

// Delete : delete file in specific path
func Delete(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

// Rotation : rotate the file
func Rotation(dir string, files []Stat, ballFileType string) error {
	if len(files) == 0 {
		return errors.New("empty files for your processing stage")
	}

	var archFiles, archPaths []string

	now := time.Now()

	dtSub := now.AddDate(0, 0, -1)
	yesterday := dtSub.Format("2006-01-02")

	for _, file := range files {
		archFile := fmt.Sprintf("%s.%s", file.FileName, yesterday)

		archPath := filepath.Join(dir, archFile)

		archFiles = append(archFiles, archFile)
		archPaths = append(archPaths, archPath)

		_ = os.Rename(file.FilePath, archPath)

		log, _ := os.OpenFile(file.FilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)

		defer log.Close()
	}

	ballFile := fmt.Sprintf("archive.%s.%s", yesterday, ballFileType)

	if ballFileType == "zip" {
		_ = tarball.Zip(dir, ballFile, archFiles)
	} else if ballFileType == "tar.gz" {
		_ = tarball.TarGz(dir, ballFile, archFiles)
	}

	for _, archPath := range archPaths {
		_ = Delete(archPath)
	}

	return nil
}

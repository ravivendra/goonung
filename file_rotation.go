package goonung

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ravivendra/goonung/rotation"
)

// FileRotation : rotate (archiving) files with day-1 and create the new one
func FileRotation(dir, fileType, ballFileType string) error {
	var stats []rotation.Stat

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		var stat rotation.Stat

		if err != nil {
			return err
		}

		ext := filepath.Ext(path)

		if ext == fmt.Sprintf(".%s", fileType) {
			stat = rotation.Stat{FilePath: path, FileName: info.Name()}

			stats = append(stats, stat)
		}

		return nil
	})

	if err != nil {
		return err
	}

	if err := rotation.Rotation(dir, stats, ballFileType); err != nil {
		return err
	}

	return nil
}

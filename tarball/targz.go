package tarball

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func addFileToTarGZ(tw *tar.Writer, in string) error {
	src, err := os.Open(in)

	if err != nil {
		return err
	}

	defer src.Close()

	info, err := src.Stat()

	if err != nil {
		return err
	}

	headers := &tar.Header{
		Name:    info.Name(),
		Size:    info.Size(),
		Mode:    int64(info.Mode()),
		ModTime: info.ModTime(),
	}

	if err := tw.WriteHeader(headers); err != nil {
		return err
	}

	if _, err := io.Copy(tw, src); err != nil {
		return err
	}

	return nil
}

// TarGz : zipping files to a tar.gz file
func TarGz(dir, out string, ins []string) error {
	dstPath := filepath.Join(dir, out)

	dst, err := os.Create(dstPath)

	if err != nil {
		return err
	}

	defer dst.Close()

	if err := dst.Chmod(os.FileMode.Perm(0600)); err != nil {
		return err
	}

	gw := gzip.NewWriter(dst)

	defer gw.Close()

	tw := tar.NewWriter(gw)

	defer tw.Close()

	for _, in := range ins {
		srcPath := filepath.Join(dir, in)

		if err := addFileToTarGZ(tw, srcPath); err != nil {
			return err
		}
	}

	return nil
}

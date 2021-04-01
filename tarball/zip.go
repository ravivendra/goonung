package tarball

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func addFileToZip(zw *zip.Writer, in string) error {
	src, err := os.Open(in)

	if err != nil {
		return err
	}

	defer src.Close()

	info, err := src.Stat()

	if err != nil {
		return err
	}

	headers, err := zip.FileInfoHeader(info)

	if err != nil {
		return err
	}

	headers.Name = info.Name()
	headers.UncompressedSize64 = uint64(info.Size())
	headers.Method = zip.Deflate

	dst, err := zw.CreateHeader(headers)

	if err != nil {
		return err
	}

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

// Zip : zipping files to a zip file
func Zip(dir, out string, ins []string) error {
	dstPath := filepath.Join(dir, out)

	dst, err := os.Create(dstPath)

	if err != nil {
		return err
	}

	defer dst.Close()

	if err := dst.Chmod(os.FileMode.Perm(0600)); err != nil {
		return err
	}

	zw := zip.NewWriter(dst)

	defer zw.Close()

	for _, in := range ins {
		srcPath := filepath.Join(dir, in)

		if err := addFileToZip(zw, srcPath); err != nil {
			return err
		}
	}

	return nil
}

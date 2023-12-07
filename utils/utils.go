package utils

import (
	"os"
	"path/filepath"
)

func GetFolderSize(folderPath string) (float64, error) {
	var size int64

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			size += info.Size()
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return BytesToMegabytes(size), nil
}

func BytesToMegabytes(bytes int64) float64 {
	const megabyte = 1024 * 1024

	return float64(bytes) / float64(megabyte)
}

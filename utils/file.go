package utils

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func VeriyStorageAndRetrievePath(name string, size int64) (string, error) {
	filename := strings.ReplaceAll(strings.Trim(name, " "), " ", "")
	fileSize := BytesToMegabytes(size)

	storagePath := os.Getenv("STORAGE")
	storageSize, _ := strconv.Atoi(os.Getenv("STORAGE_SIZE"))

	dicSize, sizeErr := GetFolderSize(storagePath)
	if sizeErr != nil {
		return "", errors.New("unable to get dictory size")
	}

	if dicSize+fileSize > float64(storageSize) {
		return "", errors.New("something went wrong when retreiving file path")
	}

	return storagePath + filename, nil
}

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

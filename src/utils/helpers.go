package devops_scripts

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"time"
)

func GetFileHash(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func GetFileSize(filePath string) (int64, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return 0, err
	}

	return info.Size(), nil
}

func GetFilesModifiedTime(filePath string) (time.Time, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return time.Time{}, err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return time.Time{}, err
	}

	return info.ModTime(), nil
}

func GetFilesInfo(filePath string) (string, int64, time.Time, error) {
	hash, err := GetFileHash(filePath)
	if err != nil {
		return "", 0, time.Time{}, err
	}

	size, err := GetFileSize(filePath)
	if err != nil {
		return "", 0, time.Time{}, err
	}

	modTime, err := GetFilesModifiedTime(filePath)
	if err != nil {
		return "", 0, time.Time{}, err
	}

	return hash, size, modTime, nil
}
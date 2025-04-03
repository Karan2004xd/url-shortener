package utils

import (
	"errors"
	"fmt"
	"os"
)

func CheckIfFileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func ReadFile(path string) (string, error) {
	if !CheckIfFileExists(path) {
		return "", errors.New(
			fmt.Sprint("The file doesn't exists at the provided path", path))
	}

	data, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

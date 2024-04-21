package filesystem

import (
	"errors"
	"os"
)

// GetFileBytes will retrieve the content of a file as bytes
func GetFileBytes(file string) ([]byte, error) {
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		fileBytes, fileErr := os.ReadFile(file)
		if fileErr == nil {
			return fileBytes, nil
		}
		return nil, err
	}

	return nil, errors.New("file does not exist")
}

// GetFileContent will retrieve the content of a file as text
func GetFileContent(file string) (string, error) {
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		fileBytes, fileErr := os.ReadFile(file)
		if fileErr == nil {
			return string(fileBytes), nil
		}
		return "", err
	}

	return "", errors.New("file does not exist")
}

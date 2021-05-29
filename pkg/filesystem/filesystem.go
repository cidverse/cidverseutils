package filesystem

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// CreateDirectory creates a new folder if not present, ignores errors
func CreateDirectory(dir string) {
	os.MkdirAll(dir, os.ModePerm)
}

/**
 * Get the relative path in relation to the rootDirectory
 */
func GetPathRelativeToDirectory(currentDirectory string, rootDirectory string) string {
	relativePath := strings.Replace(currentDirectory, rootDirectory, "", 1)
	relativePath = strings.Replace(relativePath, "\\", "/", -1)
	relativePath = strings.Trim(relativePath, "/")

	return relativePath
}

/**
 * Get the execution directory
 */
func GetExecutionDirectory() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	return filepath.Dir(ex)
}

// GetWorkingDirectory returns the current working directory
func GetWorkingDirectory() string {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return workingDir
}

// GetProjectDirectory will try to find the project directory based on repository folders (.git)
func GetProjectDirectory() (string, error) {
	currentDirectory := GetWorkingDirectory()
	var projectDirectory = ""
	directoryParts := strings.Split(currentDirectory, string(os.PathSeparator))

	for projectDirectory == "" {
		// git repository
		if _, err := os.Stat(filepath.Join(currentDirectory, "/.git")); err == nil {
			return currentDirectory, nil
		}

		// cancel at root path
		if directoryParts[0]+"\\" == currentDirectory || currentDirectory == "/" {
			return "", errors.New("didn't find any repositories for the current working directory")
		}

		currentDirectory = filepath.Dir(currentDirectory)
	}

	return "", errors.New("didn't find any repositories for the current working directory")
}

// FindFilesInDirectory will return all files with a specific extension in a directory
func FindFilesInDirectory(directory string, extension string) ([]string, error) {
	var files []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if len(extension) > 0 {
			if strings.HasSuffix(path, extension) {
				files = append(files, path)
			}
		} else {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

// CreateFileWithContent will create a new file with content
func CreateFileWithContent(file string, data string) error {
	err := ioutil.WriteFile(file, []byte(data), 0)

	if err != nil {
		return err
	}

	return nil
}

// RemoveFile will delete a file
func RemoveFile(file string) error {
	err := os.Remove(file)
	if err != nil {
		return err
	}

	return nil
}

// MoveFile will move the file to a new location
func MoveFile(oldLocation string, newLocation string) error {
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		return err
	}

	return nil
}

// GetFileBytes will retrieve the content of a file as bytes
func GetFileBytes(file string) ([]byte, error) {
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		fileBytes, fileErr := ioutil.ReadFile(file)
		if fileErr == nil {
			return fileBytes, nil
		} else {
			return nil, err
		}
	}

	return nil, errors.New("file does not exist")
}

// GetFileContent will retrieve the content of a file as text
func GetFileContent(file string) (string, error) {
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		fileBytes, fileErr := ioutil.ReadFile(file)
		if fileErr == nil {
			return string(fileBytes), nil
		} else {
			return "", err
		}
	}

	return "", errors.New("file does not exist")
}

// SaveFileContent will save a file with the provided content
func SaveFileContent(file string, content string) error {
	data := []byte(content)

	err := ioutil.WriteFile(file, data, 0)

	return err
}

// FileExists checks if the file exists and returns a boolean
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

// FileContainsString will check if a file contains the string
func FileContainsString(file string, str string) bool {
	content, contentErr := GetFileContent(file)
	if contentErr != nil {
		return false
	}

	if strings.Contains(content, str) {
		return true
	}

	return false
}

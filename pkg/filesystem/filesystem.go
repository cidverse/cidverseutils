package filesystem

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// CreateDirectory creates a new folder if not present, ignores errors
func CreateDirectory(dir string) {
	_ = os.MkdirAll(dir, os.ModePerm)
}

// GetPathRelativeToDirectory returns the relative path in relation to the rootDirectory
func GetPathRelativeToDirectory(currentDirectory string, rootDirectory string) string {
	relativePath := strings.Replace(currentDirectory, rootDirectory, "", 1)
	relativePath = strings.ReplaceAll(relativePath, "\\", "/")
	relativePath = strings.Trim(relativePath, "/")

	return relativePath
}

// GetExecutionDirectory returns the execution directory
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

// FindFilesByExtension will return all files with a specific extension in a directory
func FindFilesByExtension(directory string, extensions []string) ([]string, error) {
	var files []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			if len(extensions) > 0 {
				for _, ext := range extensions {
					if strings.HasSuffix(path, ext) {
						files = append(files, path)
						break
					}
				}
			} else {
				files = append(files, path)
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
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

// SaveFileText will save a file with the provided content
func SaveFileText(file string, content string) error {
	data := []byte(content)
	err := os.WriteFile(file, data, os.ModePerm)
	return err
}

// FileExists checks if the file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

// DirectoryExists checks if the dir exists
func DirectoryExists(dir string) bool {
	info, err := os.Stat(dir)
	if err != nil {
		return false
	}

	return info.IsDir()
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

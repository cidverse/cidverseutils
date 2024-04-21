package filesystem

import (
	"errors"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

// GetPathRelativeToDirectory returns the relative path in relation to the rootDirectory
func GetPathRelativeToDirectory(currentDirectory string, rootDirectory string) string {
	relativePath := strings.Replace(currentDirectory, rootDirectory, "", 1)
	relativePath = strings.ReplaceAll(relativePath, "\\", "/")
	relativePath = strings.Trim(relativePath, "/")

	return relativePath
}

// ExecutableDir returns the directory of the current executable
func ExecutableDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	return filepath.Dir(ex)
}

// WorkingDirOrPanic returns the current working directory or panics
func WorkingDirOrPanic() string {
	workingDir, err := os.Getwd()
	if err != nil {
		slog.Error("failed to determinate working directory", err)
		panic(err)
	}

	return workingDir
}

// GetProjectDirectory will try to find the project directory based on repository folders (.git)
func GetProjectDirectory() (string, error) {
	currentDirectory := WorkingDirOrPanic()
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

// SaveFileText will save a file with the provided content
func SaveFileText(file string, content string) error {
	data := []byte(content)
	err := os.WriteFile(file, data, os.ModePerm)
	return err
}

// CopyFile copies a file
func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return os.ErrDeadlineExceeded
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	err = out.Sync()
	return nil
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

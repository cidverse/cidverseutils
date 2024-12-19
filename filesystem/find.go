package filesystem

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/charlievieth/fastwalk"
)

type IgnoreFunc func(absPath string, name string) bool

func defaultIgnoreFunc(absPath string, name string) bool {
	return name == ".git"
}

type FilterFunc func(absPath string, name string) bool

func defaultFilterFunc(absPath string, name string) bool {
	return true
}

// FindFiles will return all files in a directory
// ignoreFunc can be used to skip directories or files
// filterFunc is used to evaluate if a file should be included in the result
func FindFiles(rootPath string, ignore IgnoreFunc, filter FilterFunc) ([]string, error) {
	var files []string

	conf := fastwalk.Config{
		Follow: false,
	}
	err := fastwalk.Walk(&conf, rootPath, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// skip if ignoreFunc returns true
		if ignore(path, info.Name()) {
			if info.IsDir() {
				return filepath.SkipDir
			}

			return nil
		}

		// add file if filterFunc returns true
		if info.IsDir() == false && filter(path, info.Name()) {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

// FindFilesByExtension will return all files with a specific extension in a directory
func FindFilesByExtension(rootPath string, extensions []string) ([]string, error) {
	return FindFiles(rootPath, defaultIgnoreFunc, func(absPath string, name string) bool {
		for _, ext := range extensions {
			if strings.HasSuffix(name, ext) {
				return true
			}
		}
		return false
	})
}

// GenerateFileMapByExtension will return a map of files by extension
// Supports one dot in the file name, see GenerateFileMapByDeepExtension for multiple dots
// Example: {"go": ["file1.go", "file2.go"], "txt": ["file1.txt"], "gz": ["file1.tar.gz"]}
func GenerateFileMapByExtension(files []string) map[string][]string {
	extensionMap := make(map[string][]string)
	for _, file := range files {
		ext := filepath.Ext(file)
		ext = strings.TrimPrefix(ext, ".")

		extensionMap[ext] = append(extensionMap[ext], file)
	}
	return extensionMap
}

// GenerateFileMapByDeepExtension will return a map of files by extension
// Supports multiple dots in the file name, see GenerateFileMapByExtension for one dot
// Example: {"go": ["file1.go", "file2.go"], "txt": ["file1.txt"], "tar.gz": ["file1.tar.gz"]}
func GenerateFileMapByDeepExtension(files []string) map[string][]string {
	extensionMap := make(map[string][]string)
	for _, file := range files {
		fileName := filepath.Base(file)
		ext := fileName

		if !strings.Contains(fileName, ".") {
			extensionMap[""] = append(extensionMap[""], file)
			continue
		}

		for {
			extIndex := strings.Index(ext, ".")
			if extIndex == -1 {
				break
			}
			ext = ext[extIndex+1:]
			extensionMap[ext] = append(extensionMap[ext], file)
		}
	}
	return extensionMap
}

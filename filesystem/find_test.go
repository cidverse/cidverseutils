package filesystem

import (
	"reflect"
	"testing"
)

func TestGenerateFileMapByExtension(t *testing.T) {
	files := []string{"file1.txt", "file2.jpg", "file3.txt", "file4.exe", "file5"}
	expected := map[string][]string{
		"txt": {"file1.txt", "file3.txt"},
		"jpg": {"file2.jpg"},
		"exe": {"file4.exe"},
		"":    {"file5"},
	}
	result := GenerateFileMapByExtension(files)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestGenerateFileMapByExtensionM(t *testing.T) {
	files := []string{"file1.txt", "file2.jpg", "file3.txt", "file4.exe", "file5", "file6.tar.gz", "file7."}
	expected := map[string][]string{
		"":       {"file5", "file7."},
		"txt":    {"file1.txt", "file3.txt"},
		"jpg":    {"file2.jpg"},
		"exe":    {"file4.exe"},
		"tar.gz": {"file6.tar.gz"},
	}
	result := GenerateFileMapByExtensionM(files)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

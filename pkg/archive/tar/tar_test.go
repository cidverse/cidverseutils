package tar

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTARCreateAndExtract(t *testing.T) {
	// Create a temporary directory to use as the input directory
	inputDir, err := os.MkdirTemp("", "test-input")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(inputDir)

	// Create a temporary directory to use as the output directory
	outputDir, err := os.MkdirTemp("", "test-output")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(outputDir)

	// Create a test file in the input directory
	testFile := filepath.Join(inputDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("hello world"), 0644); err != nil {
		t.Fatal(err)
	}

	// Create a tar archive of the input directory
	archiveFile := filepath.Join(outputDir, "test.tar")
	if err := TARCreate(inputDir, archiveFile); err != nil {
		t.Fatal(err)
	}

	// Extract the tar archive into the output directory
	if err := TARExtract(archiveFile, outputDir); err != nil {
		t.Fatal(err)
	}

	// Check that the extracted file has the same contents as the original file
	extractedFile := filepath.Join(outputDir, "test.txt")
	extractedData, err := os.ReadFile(extractedFile)
	if err != nil {
		t.Fatal(err)
	}

	expectedData, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatal(err)
	}

	if string(extractedData) != string(expectedData) {
		t.Fatalf("unexpected file contents: %s", string(extractedData))
	}
}

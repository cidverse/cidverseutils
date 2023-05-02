package zip

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateAndExtract(t *testing.T) {
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

	// Create a test file to use as input
	inputFile := filepath.Join(inputDir, "test.txt")
	if err := os.WriteFile(inputFile, []byte("test"), 0644); err != nil {
		t.Fatal(err)
	}

	// Create a test zip archive
	archiveFile := filepath.Join(outputDir, "test.zip")
	if err := Create(inputDir, archiveFile); err != nil {
		t.Fatal(err)
	}

	// Extract the test zip archive
	if err := Extract(archiveFile, outputDir); err != nil {
		t.Fatal(err)
	}

	// Check that the extracted file matches the original file
	extractedFile := filepath.Join(outputDir, "test.txt")
	extractedData, err := os.ReadFile(extractedFile)
	if err != nil {
		t.Fatal(err)
	}
	if string(extractedData) != "test" {
		t.Errorf("extracted file does not match input file")
	}
}

package compress

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// ZIPCreate creates a zip archive of the directory at the given path.
func ZIPCreate(inputDirectory string, outputFile string) error {
	// Create a new zip file to write to
	newZipFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	// Create a new zip archive
	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Walk the directory to get all files and subdirectories
	return filepath.Walk(inputDirectory, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// If file is a directory, skip it
		if fileInfo.IsDir() {
			return nil
		}

		// Open the file for reading
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// Create a new file header for the file
		header, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			return err
		}

		// Set the name of the file header to include the path
		header.Name = filePath[len(inputDirectory)+1:]

		// Add the file header to the zip archive
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// Write the contents of the file to the zip archive
		if _, err = io.Copy(writer, file); err != nil {
			return err
		}

		return nil
	})
}

// ZIPExtract unzips the zip archive at the given path into the given directory.
func ZIPExtract(archiveFile string, outputDirectory string) error {
	// Open the zip archive for reading
	zipReader, err := zip.OpenReader(archiveFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	// Create the directory to extract files into if it doesn't exist
	if err = os.MkdirAll(outputDirectory, os.ModePerm); err != nil {
		return err
	}

	// Extract all files from the zip archive
	for _, file := range zipReader.File {
		// Construct the destination path for the extracted file
		destPath := filepath.Join(outputDirectory, file.Name)

		// Create any necessary directories in the destination path
		if file.FileInfo().IsDir() {
			if err = os.MkdirAll(destPath, os.ModePerm); err != nil {
				return err
			}
			continue
		} else {
			if err = os.MkdirAll(filepath.Dir(destPath), os.ModePerm); err != nil {
				return err
			}
		}

		// Open the file from the zip archive for reading
		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		// Create a new file in the destination path for writing
		newFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer newFile.Close()

		// Copy the contents of the file from the zip archive to the new file
		if _, err := io.Copy(newFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

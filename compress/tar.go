package compress

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

// TARCreate creates a tar archive of the directory at the given path.
func TARCreate(inputDirectory string, outputFile string) error {
	// Create a new tar file to write to
	newTarFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer newTarFile.Close()

	// Create a new tar writer
	tarWriter := tar.NewWriter(newTarFile)
	defer tarWriter.Close()

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

		// Create a new tar header for the file
		header := &tar.Header{
			Name: filePath[len(inputDirectory)+1:],
			Mode: int64(fileInfo.Mode().Perm()),
			Size: fileInfo.Size(),
		}

		// Write the header to the tar archive
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		// Write the contents of the file to the tar archive
		if _, err := io.Copy(tarWriter, file); err != nil {
			return err
		}

		return nil
	})
}

// TARExtract extracts a tar archive at the given path into the given directory.
func TARExtract(archiveFile string, outputDirectory string) error {
	// Open the tar archive for reading
	tarFile, err := os.Open(archiveFile)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	// Create a new tar reader
	tarReader := tar.NewReader(tarFile)

	// Extract all files from the tar archive
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Construct the destination path for the extracted file
		destPath := filepath.Join(outputDirectory, header.Name)

		// Create any necessary directories in the destination path
		if header.FileInfo().IsDir() {
			if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
				return err
			}
			continue
		} else {
			if err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm); err != nil {
				return err
			}
		}

		// Create a new file in the destination path for writing
		newFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer newFile.Close()

		// Copy the contents of the file from the tar archive to the new file
		if _, err := io.Copy(newFile, tarReader); err != nil {
			return err
		}
	}

	return nil
}

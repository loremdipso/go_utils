package go_utils

import (
	"fmt"
	"io"
	"os"

	"github.com/cloudfoundry/bytefmt"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func GetFileSize(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		return -1
	}

	return info.Size()
}

func HumanReadableSize(size int64) string {
	return bytefmt.ByteSize(uint64(size))
}

func MoveFile(oldLocation, newLocation string) error {
	// first, try a simple rename
	err := os.Rename(oldLocation, newLocation)
	if err != nil {

		// failing that, assume it's because the file's on another device
		return actuallyMoveFile(oldLocation, newLocation)
	}

	return nil
}

func actuallyMoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}
	return nil
}

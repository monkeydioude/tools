package tools

import (
	"os"
	"testing"
)

func TestDestFileExistsAndIsWritable(t *testing.T) {
	err := HandleDestinationFile("testdata/readable-and-writable-file.json")
	if err != nil {
		t.Error("File should exist and be readable")
	}
}

func TestIFailOnNotWritableDestFile(t *testing.T) {
	file := "testdata/not-writable-file.json"
	// Read only
	os.Chmod(file, 0466)

	err := HandleDestinationFile(file)
	if err == nil {
		t.Error("File should exist and should not be writable")
	}

	os.Chmod(file, 0666)
}

func TestIJustReturnWhenDestFileDoesNotExist(t *testing.T) {
	err := HandleDestinationFile("testdata/unknownsource")
	if err.Error() != "[ERR ] File does not exist" {
		t.Error("File should not exist")
	}
}

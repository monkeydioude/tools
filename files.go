package tools

import (
	"os"
)

type FileError struct {
	err     string
	ErrMask int
}

const (
	FILE_NEXIST = 1 << iota
	FILE_NWRITE
)

var ErrLabel = map[int]string{
	FILE_NEXIST: "File does not exist",
	FILE_NWRITE: "File exists but can not be written on",
}

// FileExists verifies if file exists/doesn't and is readable
func FileExists(p string) *FileError {
	f, _ := os.Stat(p)

	if f == nil {
		return &FileError{
			err:     ErrLabel[FILE_NEXIST],
			ErrMask: FILE_NEXIST,
		}
	}

	if f.Mode()&0200 == 0 {
		return &FileError{
			err:     ErrLabel[FILE_NWRITE],
			ErrMask: FILE_NWRITE,
		}
	}

	return nil
}

func (e *FileError) Error() string {
	return e.err
}

func (e *FileError) String() string {
	return e.Error()
}

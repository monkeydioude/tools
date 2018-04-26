package tools

import (
	"errors"
	"os"
)

// Verify does not exist or exists and is readable
func HandleDestinationFile(p string) error {
	f, _ := os.Stat(p)

	if f == nil {
		return errors.New("[ERR ] File does not exist")
	}

	if f.Mode()&0200 == 0 {
		return errors.New("[ERR ] File exists but can not be written")
	}

	return nil
}

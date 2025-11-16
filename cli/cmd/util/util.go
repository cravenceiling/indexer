package util

import "os"

// CheckEmpty returns true if the file with path filePath is empty.
// And false otherwise, it returns true, err if an error was returned
// when getting the file information.
func CheckEmpty(filePath string) (bool, error) {
	fi, err := os.Stat(filePath)
	if err != nil {
		return true, err
	}

	return (fi.Size() == 0), nil
}

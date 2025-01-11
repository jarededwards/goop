package utils

import (
	"fmt"
	"os"
)

func CreateDirIfNotExist(dir string) error {
	if err := os.MkdirAll(dir, 0o777); err != nil {
		return fmt.Errorf("unable to create directory %q: %w", dir, err)
	}

	return nil
}

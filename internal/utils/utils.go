package utils

import (
	"fmt"
	"os"
	"strings"
)

func CreateDirIfNotExist(dir string) error {
	if err := os.MkdirAll(dir, 0o777); err != nil {
		return fmt.Errorf("unable to create directory %q: %w", dir, err)
	}

	return nil
}

func DeleteNestedKey(m map[string]interface{}, path string) {
	parts := strings.Split(path, ".")

	// Navigate to the parent object
	current := m
	for i := 0; i < len(parts)-1; i++ {
		if next, ok := current[parts[i]].(map[string]interface{}); ok {
			current = next
		} else {
			// Path doesn't exist, nothing to delete
			return
		}
	}

	// Delete the final key
	delete(current, parts[len(parts)-1])
}

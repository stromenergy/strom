package util

import (
	"encoding/base64"
	"errors"
	"os"
)

func ResolveFromFile(base64Value, filePath string) ([]byte, error) {
	if len(filePath) > 0 {
		fileBytes, err := os.ReadFile(filePath)

		if err != nil {
			LogError("STR071: Error reading file data", err)
			return nil, errors.New("Error reading file data")

		}

		return fileBytes, nil
	}

	value, err := base64.StdEncoding.DecodeString(base64Value)

	if err != nil {
		LogError("STR072: Error decoding base64", err)
		return nil, errors.New("Error decoding base64")
	}

	return value, nil
}

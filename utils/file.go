package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetFileExtension(platform string) string {
	switch platform {
	case "android":
		return "xml"
	case "ios":
		return "strings"
	default:
		return "txt"
	}
}

func GetPlatformFolderName(platform string) string {
	switch platform {
	case "android":
		return "Android"
	case "ios":
		return "iOS"
	default:
		return "unknown"
	}
}

func HasTranslation() bool {
	if _, err := os.Stat("../../output/Android"); os.IsNotExist(err) {
		// path does not exist
		return false
	}
	return true
}

func CreateDirectoryIfNotExist(path string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}
	}
	return nil
}

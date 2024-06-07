package utils

import (
	"path/filepath"
	"strings"
)

func IsImageFile(filename string) bool {
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}

	ext := strings.ToLower(filepath.Ext(filename))

	return allowedExtensions[ext]
}

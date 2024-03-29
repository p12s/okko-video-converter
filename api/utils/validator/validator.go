package validator

import (
	"github.com/spf13/viper"
)

func IsFileTypeValid(fileType string) bool {
	switch fileType {
	case
		"video/quicktime",
		"video/mp4",
		"video/3gpp",
		"video/avi",
		"video/x-flv",
		"application/octet-stream",
		"video/x-msvideo",
		"video/mpeg":
		return true
	}
	return false
}

func IsFileSizeExceeded(fileSize int64) bool {
	return fileSize <= viper.GetInt64("maxFileSizeInMb")*1024
}

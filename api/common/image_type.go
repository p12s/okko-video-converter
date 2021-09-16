package common

import (
	"path/filepath"
	"strings"
)

type ImageType int

const (
	UNKNOWN ImageType = iota
	MP4
	AVI
	MPEG
	MOV
	FLV
	WEBM
	MKV
)

func GetFileType(fileName string) ImageType {
	ext := filepath.Ext(strings.ToLower(fileName))

	switch ext {
	case ".mp4":
		return MP4
	case ".avi":
		return AVI
	case ".mpeg":
		return MPEG
	case ".mov":
		return MOV
	case ".flv":
		return FLV
	case ".webm":
		return WEBM
	case ".mkv":
		return MKV
	default:
		return UNKNOWN
	}
}

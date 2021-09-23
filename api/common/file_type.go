package common

import (
	"path/filepath"
	"strings"
)

type FileType int

const (
	UNKNOWN FileType = iota
	MP4
	AVI
	MPEG
	MOV
	FLV
	WEBM
	MKV
)

func GetFileType(fileName string) FileType {
	// добавляем точку в начале на случай, если придет строка вида "flv".
	// без точки filepath.Ext не распознает расширение
	ext := filepath.Ext(strings.ToLower("." + fileName))

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

func GetFileExt(fileType FileType) string {
	switch fileType {
	case MP4:
		return "mp4"
	case AVI:
		return "avi"
	case MPEG:
		return "mpeg"
	case MOV:
		return "mov"
	case FLV:
		return "flv"
	case WEBM:
		return "webm"
	case MKV:
		return "mkv"
	default:
		return ""
	}
}

package common

import (
	"errors"
	"github.com/google/uuid"
)

const (
	EVENT_VIDEO_CONVERT = "video.convert"
)

type User struct {
	Id   int       `json:"-" db:"id"`
	Code uuid.UUID `json:"code" db:"code" binding:"required"`
}

type File struct {
	Id           int           `json:"-" db:"id"`
	Path         string        `json:"path" db:"path"`
	Name         string        `json:"name" db:"name"`
	UserId       int           `json:"-" db:"user_id"`
	KiloByteSize int64         `json:"kilo_byte_size" db:"kilo_byte_size"`
	PrevImage    string        `json:"prev_image" db:"prev_image"`
	Status       ProcessStatus `json:"status" db:"status"`
	ErrorMessage string        `json:"error_message" db:"error_message"`
}

type UploadedFile struct {
	Path         string   `json:"path,omitempty"`
	Name         string   `json:"name,omitempty"`
	KiloByteSize int64    `json:"kilo_byte_size,omitempty"`
	Error        string   `json:"error"`
	IsError      bool     `json:"is_error"`
	Type         FileType `json:"file_type",omitempty`
	PrevImage    string   `json:"prev_image"`
}

type ResizedImg struct {
	Path string `json:"-"`
	Name string `json:"-"`
}

type SizeCoefficient struct {
	Coef float64
	W    uint
	H    uint
}

type Width struct {
	Width      int
	SizeFields []SizeCoefficient
}

type WidthList struct {
	List  []Width
	Total int
}

type ResizeProcess struct {
	Total   int
	Current int
}

type VideoConvertData struct { // если будет несколько типов данных - заменить интерфейсом
	UserCode     string   `json:"user_code"`
	Path         string   `json:"path"`
	TargetFormat FileType `json:"target_format"`
}

func (t VideoConvertData) Validate() error {
	if t.UserCode == "" {
		return errors.New("userCode is required")
	}
	if t.Path != "" {
		return errors.New("path is required")
	}
	if t.TargetFormat != UNKNOWN {
		return errors.New("targetFormat is required")
	}
	return nil
}

type Event struct {
	Type  string // может пригодиться для разделения видов событий, но пока будет только 1
	Value VideoConvertData
}

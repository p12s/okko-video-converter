package common

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id   int       `json:"-" db:"id"`
	Code uuid.UUID `json:"code" db:"code" binding:"required"`
}

type File struct {
	Id           int    `json:"-" db:"id"`
	Path         string `json:"path" db:"path"`
	Name         string `json:"name" db:"name"`
	UserId       int    `json:"-" db:"user_id"`
	KiloByteSize int64  `json:"kilo_byte_size" db:"kilo_byte_size"`
	PrevImage    string `json:"prev_image" db:"prev_image"`
}

type ResizeOptions struct {
	Id           int          `json:"-" db:"id"`
	UserId       int          `json:"user_id" db:"user_id"`
	Options      string       `json:"options" db:"options"`
	StartDate    time.Time    `json:"start_date" db:"start_date"`
	FinishDate   time.Time    `json:"finish_date" db:"finish_date"`
	Status       ResizeStatus `json:"status" db:"status"`
	TotalCount   int          `json:"total_count" db:"total_count"`
	Current      int          `json:"current" db:"current"`
	ErrorMessage string       `json:"error_message" db:"error_message"`
}

type UploadedFile struct {
	Path         string    `json:"path,omitempty"`
	Name         string    `json:"name,omitempty"`
	KiloByteSize int64     `json:"kilo_byte_size,omitempty"`
	Error        string    `json:"error"`
	IsError      bool      `json:"is_error"`
	Type         ImageType `json:"file_type",omitempty`
	PrevImage    string    `json:"prev_image"`
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

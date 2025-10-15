package model

import (
	"gorm.io/gorm"
)

// FileRecord stores metadata of uploaded file
type FileRecord struct {
	gorm.Model

	PublicID  string    `json:"public_id"`
	URL       string    `json:"url"`
	Filename  string    `json:"filename"`
	Bytes     int64     `json:"bytes"`
}

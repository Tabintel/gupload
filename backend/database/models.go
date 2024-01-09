// backend/database/models.go

package database

import "github.com/jinzhu/gorm"

// Image represents the model for storing image information in the database.
type Image struct {
	gorm.Model
	OriginalURL string `gorm:"not null"`
	ThumbnailURL string `gorm:"not null"`
}

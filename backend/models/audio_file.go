package models

import "time"

type AudioFile struct {
	ID        uint `gorm:"primaryKey"`
	FileName  string
	Format    string
	CreatedAt time.Time
}

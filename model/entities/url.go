package entities

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	OriginalUrl string `gorm:"type:varchar(255);not null"`
	ShortUrl    string `gorm:"type:varchar(10);uniqueIndex"`
}

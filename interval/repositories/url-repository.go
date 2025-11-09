package repositories

import (
	"gorm.io/gorm"
	"sdemir.com/model/entities"
)

type UrlRepository interface {
	SaveUrl(request *entities.Url) (*entities.Url, error)
	GetUrl(shortUrl string) (*entities.Url, error)
}

type urlRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UrlRepository {
	return &urlRepository{db: db}
}

func (r *urlRepository) SaveUrl(request *entities.Url) (*entities.Url, error) {
	result := r.db.Save(request)
	if result.Error != nil {
		return nil, result.Error
	}
	return request, nil
}

func (r *urlRepository) GetUrl(shortUrl string) (*entities.Url, error) {
	url := &entities.Url{}
	result := r.db.Where("short_url = ?", shortUrl).First(url)
	if result.Error != nil {
		return nil, result.Error
	}
	return url, nil
}

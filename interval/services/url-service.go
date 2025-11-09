package services

import (
	"errors"

	"sdemir.com/interval/repositories"
	"sdemir.com/model/entities"
	"sdemir.com/utils"
)

type UrlService interface {
	CreateShortUrl(request string) (*entities.Url, error)
	GetUrl(request string) (*entities.Url, error)
}

type urlService struct {
	repo repositories.UrlRepository
}

func NewUrlService(repo repositories.UrlRepository) UrlService {
	return &urlService{repo: repo}
}

func (s urlService) CreateShortUrl(request string) (*entities.Url, error) {
	if request == "" {
		return nil, errors.New("url boş olamaz")
	}
	shortUrl := utils.GenerateShortCode(6)
	newUrl := &entities.Url{}
	newUrl.OriginalUrl = request
	newUrl.ShortUrl = shortUrl
	result, err := s.repo.SaveUrl(newUrl)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return result, nil
}

func (s urlService) GetUrl(request string) (*entities.Url, error) {
	if request == "" {
		return nil, errors.New("url boş olamaz")
	}
	foundUrl, err := s.repo.GetUrl(request)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return foundUrl, nil
}

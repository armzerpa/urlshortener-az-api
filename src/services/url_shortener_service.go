package services

import (
	"github.com/armzerpa/urlshortener-az-api/src/domain"
	"github.com/armzerpa/urlshortener-az-api/src/repository/db"
	"github.com/armzerpa/urlshortener-az-api/src/utils/errors"
)

type ShortenerService interface {
	SaveUrl(string) (*domain.Url, errors.RestError)
	GetUrl(string) (*domain.Url, errors.RestError)
	GetById(string) (*domain.Url, errors.RestError)
}

type shortener struct {
	dbRepo db.UrlRepository
}

func NewShortenerService(repo db.UrlRepository) ShortenerService {
	return &shortener{
		dbRepo: repo,
	}
}

func (s *shortener) SaveUrl(longUrl string) (*domain.Url, errors.RestError) {
	if len(longUrl) == 0 {
		return nil, errors.NewBadRequestError("url is empty")
	}

	url, err := s.dbRepo.Save(longUrl)
	if err != nil {
		return nil, errors.NewInternalServerError("something went wrong saving url")
	}
	return url, nil
}

func (s *shortener) GetUrl(shortUrl string) (*domain.Url, errors.RestError) {
	url, err := s.dbRepo.Get(shortUrl)
	if err != nil {
		return nil, errors.NewInternalServerError("something went wrong getting the url")
	}

	if len(url.ShortUrl) == 0 {
		return nil, errors.NewNotFoundError("no results")
	}
	return url, nil
}

func (s *shortener) GetById(id string) (*domain.Url, errors.RestError) {
	url, err := s.dbRepo.GetById(id)
	if err != nil {
		return nil, errors.NewInternalServerError("something went wrong getting the url")
	}

	if len(url.ShortUrl) == 0 {
		return nil, errors.NewNotFoundError("no results")
	}
	return url, nil
}

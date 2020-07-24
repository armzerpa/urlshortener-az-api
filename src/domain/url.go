package domain

import (
	"log"

	"github.com/armzerpa/urlshortener-az-api/src/utils/errors"
)

type Url struct {
	ID       string `json:"id,omitempty"`
	LongUrl  string `json:"longUrl,omitempty"`
	ShortUrl string `json:"shortUrl,omitempty"`
}

type ShortenerRequest struct {
	LongUrl string `json:"url"`
}

func (sr *ShortenerRequest) Validate() errors.RestError {
	log.Println("validation start")
	if len(sr.LongUrl) == 0 {
		return errors.NewBadRequestError("invalid url input")
	}
	return nil
}

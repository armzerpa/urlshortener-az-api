package db

import (
	"github.com/armzerpa/urlshortener-az-api/src/repository/clients"
	"github.com/armzerpa/urlshortener-az-api/src/utils/hash"

	"github.com/armzerpa/urlshortener-az-api/src/domain"
	"github.com/couchbase/gocb"
)

const (
	queryGetUrlDocumentWithLongUrl  = "SELECT `shortener-url`.* FROM `shortener-url` WHERE longUrl = $1"
	queryGetUrlDocumentWithShortUrl = "SELECT `shortener-url`.* FROM `shortener-url` WHERE shortUrl = $1"
)

var (
	domainUrl string
)

type UrlRepository interface {
	Save(string) (*domain.Url, error)
	Get(string) (*domain.Url, error)
	GetById(string) (*domain.Url, error)
}

type urlCouchbaseRepo struct {
}

func (u *urlCouchbaseRepo) Save(longUrl string) (*domain.Url, error) {
	var url domain.Url
	var n1qlParams []interface{}
	n1qlParams = append(n1qlParams, longUrl)
	query := gocb.NewN1qlQuery(queryGetUrlDocumentWithLongUrl)
	bucket := clients.GetBucket()
	rows, err := bucket.ExecuteN1qlQuery(query, n1qlParams)
	if err != nil {
		return nil, err
	}
	var row domain.Url
	rows.One(&row)
	if row == (domain.Url{}) {
		id := hash.GetHasher().GetHashId()
		url.ID = id
		url.ShortUrl = domainUrl + id
		url.LongUrl = longUrl
		bucket.Insert(url.ID, url, 0)
	} else {
		url = row
	}
	return &url, nil
}

func (u *urlCouchbaseRepo) Get(shortUrl string) (*domain.Url, error) {
	var n1qlParams []interface{}
	query := gocb.NewN1qlQuery(queryGetUrlDocumentWithShortUrl)
	n1qlParams = append(n1qlParams, shortUrl)
	rows, _ := clients.GetBucket().ExecuteN1qlQuery(query, n1qlParams)
	var row domain.Url
	rows.One(&row)

	return &row, nil
}

func (u *urlCouchbaseRepo) GetById(id string) (*domain.Url, error) {
	var url domain.Url
	clients.GetBucket().Get(id, &url)

	return &url, nil
}

func NewUrlRepository(d string) UrlRepository {
	domainUrl = d
	return &urlCouchbaseRepo{}
}

package app

import (
	"net/http"

	"github.com/armzerpa/urlshortener-az-api/src/repository/db"
	"github.com/armzerpa/urlshortener-az-api/src/services"

	"github.com/armzerpa/urlshortener-az-api/src/controllers"

	"github.com/gin-gonic/gin"
)

const (
	domainUrl = "http://localhost:8080/u/"
)

var (
	router = gin.Default()
)

func StartApplication() {
	urlHandler := controllers.NewHandler(services.NewShortenerService(db.NewUrlRepository(domainUrl)))

	v1 := router.Group("/v1")
	{
		v1.GET("/ping", Ping)
		v1.GET("/shortener/url", urlHandler.GetUrl)
		v1.POST("/shortener", urlHandler.CreateUrl)
	}
	router.GET("u/:url_id", urlHandler.RedirectUrl)

	router.Run(":8080")
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

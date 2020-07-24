package app

import (
	"net/http"

	"github.com/armzerpa/urlshortener-az-api/src/repository/db"
	"github.com/armzerpa/urlshortener-az-api/src/services"

	"github.com/armzerpa/urlshortener-az-api/src/controllers"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	urlHandler := controllers.NewHandler(services.NewShortenerService(db.NewUrlRepository()))
	router.GET("/ping", Ping)
	router.GET("u/:url_id", urlHandler.RedirectUrl)
	router.GET("/shortener/url", urlHandler.GetUrl)

	router.POST("/url", urlHandler.CreateUrl)

	router.Run(":8080")
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

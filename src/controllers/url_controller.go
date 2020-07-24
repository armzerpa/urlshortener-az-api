package controllers

import (
	"net/http"

	"github.com/armzerpa/urlshortener-az-api/src/domain"
	"github.com/armzerpa/urlshortener-az-api/src/services"
	"github.com/armzerpa/urlshortener-az-api/src/utils/errors"
	"github.com/gin-gonic/gin"
)

type HandlerController interface {
	CreateUrl(*gin.Context)
	GetUrl(*gin.Context)
	RedirectUrl(*gin.Context)
}

type handler struct {
	service services.ShortenerService
}

func NewHandler(service services.ShortenerService) HandlerController {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateUrl(c *gin.Context) {
	var urlRequest domain.ShortenerRequest
	error := c.BindJSON(&urlRequest)
	if error != nil {
		restErr := errors.NewBadRequestError("invalid json object")
		c.JSON(restErr.Status(), restErr)
		return
	}

	url, err := h.service.SaveUrl(urlRequest.LongUrl)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, url)
}

func (h *handler) GetUrl(c *gin.Context) {
	url, err := h.service.GetUrl(c.Query("short_url"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, url)
}

func (h *handler) RedirectUrl(c *gin.Context) {
	url, err := h.service.GetById(c.Param("url_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.Redirect(301, url.LongUrl)
}

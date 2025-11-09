package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sdemir.com/interval/services"
)

type UrlHandler interface {
	CreateShortenUrl(c *gin.Context)
	RedirectToUrl(c *gin.Context)
}

type urlHandler struct {
	urlService services.UrlService
}

type ShortenUrlRequestDTO struct {
	Url string `json:"url"`
}

func NewUrlHandler(urlService services.UrlService) UrlHandler {
	return &urlHandler{urlService: urlService}
}

func (h *urlHandler) CreateShortenUrl(c *gin.Context) {
	var request ShortenUrlRequestDTO
	c.ShouldBindJSON(&request)
	result, err := h.urlService.CreateShortUrl(request.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (h *urlHandler) RedirectToUrl(c *gin.Context) {
	var request string = c.Param("short-url")
	result, err := h.urlService.GetUrl(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.Redirect(http.StatusMovedPermanently, result.OriginalUrl)
}

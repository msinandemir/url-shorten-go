package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"sdemir.com/configurations"
	"sdemir.com/interval/handlers"
)

func NewApi(urlHandler handlers.UrlHandler) {
	APP_PORT := fmt.Sprintf(":%d", configurations.GetEnv.APP_PORT)
	app := gin.Default()
	setupUrlRoutes(app, urlHandler)

	if err := app.Run(APP_PORT); err != nil {
		fmt.Printf("Server başlatılırken hata oluştu. Error: %s", err)
	}
}

func setupUrlRoutes(engine *gin.Engine, urlHandler handlers.UrlHandler) *gin.Engine {
	engine.POST("/shorten", urlHandler.CreateShortenUrl)
	engine.GET("/:short-url", urlHandler.RedirectToUrl)
	return engine
}

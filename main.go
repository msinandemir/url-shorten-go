package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"sdemir.com/configurations"
)

func main() {
	configurations.ConnectDB()
	APP_PORT := fmt.Sprintf(":%d", configurations.GetEnv.APP_PORT)

	app := gin.Default()
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "HELLO FROM API",
		})
	})

	app.Run(APP_PORT)
}

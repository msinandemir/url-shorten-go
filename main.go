package main

import (
	"sdemir.com/cmd/api"
	"sdemir.com/configurations"
	"sdemir.com/interval/handlers"
	"sdemir.com/interval/repositories"
	"sdemir.com/interval/services"
	"sdemir.com/model/entities"
)

func main() {
	db := configurations.ConnectDB()
	db.DB.AutoMigrate(&entities.Url{})

	urlRepository := repositories.NewUserRepository(db.DB)
	urlService := services.NewUrlService(urlRepository)
	urlHandler := handlers.NewUrlHandler(urlService)
	api.NewApi(urlHandler)
}

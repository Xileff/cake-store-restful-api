package main

import (
	"felixsavero/cake-store-restful-api/app"
	"felixsavero/cake-store-restful-api/controller"
	"felixsavero/cake-store-restful-api/helper"
	"felixsavero/cake-store-restful-api/repository"
	"felixsavero/cake-store-restful-api/service"

	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validator := validator.New()

	cakeRepository := repository.NewCakeRepository()
	cakeService := service.NewCakeService(cakeRepository, db, validator)
	cakeController := controller.NewCakeController(cakeService)

	router := app.NewRouter(cakeController)

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

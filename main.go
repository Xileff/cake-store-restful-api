package main

import (
	"felixsavero/cake-store-restful-api/app"
	"felixsavero/cake-store-restful-api/controller"
	"felixsavero/cake-store-restful-api/exception"
	"felixsavero/cake-store-restful-api/helper"
	"felixsavero/cake-store-restful-api/repository"
	"felixsavero/cake-store-restful-api/service"

	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validator := validator.New()

	cakeRepository := repository.NewCakeRepository()
	cakeService := service.NewCakeService(cakeRepository, db, validator)
	cakeController := controller.NewCakeController(cakeService)

	router := httprouter.New()

	router.GET("/cakes", cakeController.FindAll)
	router.GET("/cakes/:id", cakeController.FindById)
	router.POST("/cakes", cakeController.Create)
	router.PUT("/cakes/:id", cakeController.Update)
	router.DELETE("/cakes/:id", cakeController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

package app

import (
	"felixsavero/cake-store-restful-api/controller"
	"felixsavero/cake-store-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(cakeController controller.CakeController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/cakes", cakeController.FindAll)
	router.GET("/cakes/:id", cakeController.FindById)
	router.POST("/cakes", cakeController.Create)
	router.PUT("/cakes/:id", cakeController.Update)
	router.DELETE("/cakes/:id", cakeController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}

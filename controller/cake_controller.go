package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CakeController interface {
	Create(writer http.ResponseWriter, request http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request http.Request, params httprouter.Params)
	Find(writer http.ResponseWriter, request http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request http.Request, params httprouter.Params)
}

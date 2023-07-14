package exception

import (
	"felixsavero/cake-store-restful-api/helper"
	"felixsavero/cake-store-restful-api/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	_, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")

		response := web.MessageResponse{
			Status:  "fail",
			Message: "Oops, cake not found",
		}

		helper.WriteToResponseBody(writer, response, http.StatusNotFound)
		return true
	}

	return false
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	_, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")

		response := web.MessageResponse{
			Status:  "fail",
			Message: "Oops, all data are required.",
		}

		helper.WriteToResponseBody(writer, response, http.StatusBadRequest)
		return true
	}

	return false
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")

	response := web.MessageResponse{
		Status:  "fail",
		Message: "Oops, something went wrong on our side",
	}

	helper.WriteToResponseBody(writer, response, http.StatusInternalServerError)
}

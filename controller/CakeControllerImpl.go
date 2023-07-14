package controller

import (
	"felixsavero/cake-store-restful-api/helper"
	"felixsavero/cake-store-restful-api/model/web"
	"felixsavero/cake-store-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CakeControllerImpl struct {
	CakeService service.CakeService
}

func NewCakeController(cakeService service.CakeService) CakeController {
	return &CakeControllerImpl{
		CakeService: cakeService,
	}
}

func (controller *CakeControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cakeCreateRequest := web.CakeCreateRequest{}
	helper.ReadFromRequestBody(request, &cakeCreateRequest)

	cakeResponse := controller.CakeService.Create(request.Context(), cakeCreateRequest)
	webResponse := web.WebResponse{
		Code:   201,
		Status: "OK",
		Data:   cakeResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CakeControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cakeUpdateRequest := web.CakeUpdateRequest{}
	helper.ReadFromRequestBody(request, &cakeUpdateRequest)

	cakeId := params.ByName("id")
	id, err := strconv.Atoi(cakeId)
	helper.PanicIfError(err)

	cakeUpdateRequest.Id = id

	cakeResponse := controller.CakeService.Update(request.Context(), cakeUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cakeResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CakeControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cakeId := params.ByName("id")
	id, err := strconv.Atoi(cakeId)
	helper.PanicIfError(err)

	controller.CakeService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CakeControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cakeId := params.ByName("id")
	id, err := strconv.Atoi(cakeId)
	helper.PanicIfError(err)

	cakeResponse := controller.CakeService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cakeResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CakeControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cakesResponse := controller.CakeService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cakesResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

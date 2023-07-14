package service

import (
	"context"
	"database/sql"
	"felixsavero/cake-store-restful-api/helper"
	"felixsavero/cake-store-restful-api/model/domain"
	"felixsavero/cake-store-restful-api/model/web"
	"felixsavero/cake-store-restful-api/repository"

	"github.com/go-playground/validator/v10"
)

type CakeServiceImpl struct {
	CakeRepository repository.CakeRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewCakeService(cakeRepository repository.CakeRepository, DB *sql.DB, validate *validator.Validate) CakeService {
	return &CakeServiceImpl{
		CakeRepository: cakeRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *CakeServiceImpl) Create(ctx context.Context, request web.CakeCreateRequest) web.CakeResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cake := domain.Cake{
		Title:       request.Title,
		Description: request.Description,
		Rating:      request.Rating,
		Image:       request.Image,
	}

	cake = service.CakeRepository.Save(ctx, tx, cake)

	return helper.ToCakeResponse(cake)
}

func (service *CakeServiceImpl) Update(ctx context.Context, request web.CakeUpdateRequest) web.CakeResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cake, err := service.CakeRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	cake.Title = request.Title
	cake.Description = request.Description
	cake.Rating = request.Rating
	cake.Image = request.Image

	cake = service.CakeRepository.Update(ctx, tx, cake)

	return helper.ToCakeResponse(cake)
}

func (service *CakeServiceImpl) Delete(ctx context.Context, cakeId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cake, err := service.CakeRepository.FindById(ctx, tx, cakeId)
	helper.PanicIfError(err)

	service.CakeRepository.Delete(ctx, tx, cake)
}

func (service *CakeServiceImpl) FindById(ctx context.Context, cakeId int) web.CakeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cake, err := service.CakeRepository.FindById(ctx, tx, cakeId)
	helper.PanicIfError(err)

	return helper.ToCakeResponse(cake)
}

func (service *CakeServiceImpl) FindAll(ctx context.Context) []web.CakeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cakes := service.CakeRepository.FindAll(ctx, tx)

	return helper.ToCakeResponses(cakes)
}

package service

import (
	"context"
	"felixsavero/cake-store-restful-api/model/web"
)

type CakeService interface {
	Create(ctx context.Context, request web.CakeCreateRequest) web.CakeResponse
	Update(ctx context.Context, request web.CakeUpdateRequest) web.CakeResponse
	Delete(ctx context.Context, cakeId int)
	FindById(ctx context.Context, cakeId int) web.CakeResponse
	FindAll(ctx context.Context) []web.CakeResponse
}

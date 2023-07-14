package helper

import (
	"felixsavero/cake-store-restful-api/model/domain"
	"felixsavero/cake-store-restful-api/model/web"
)

func ToCakeResponse(cake domain.Cake) web.CakeResponse {
	return web.CakeResponse{
		Id:          cake.Id,
		Title:       cake.Title,
		Description: cake.Description,
		Rating:      cake.Rating,
		Image:       cake.Image,
		CreatedAt:   cake.CreatedAt,
		UpdatedAt:   cake.UpdatedAt,
	}
}

func ToCakeResponses(cakes []domain.Cake) []web.CakeResponse {
	var cakeResponses []web.CakeResponse
	for _, cake := range cakes {
		cakeResponses = append(cakeResponses, ToCakeResponse(cake))
	}

	return cakeResponses
}

package web

type CakeCreateRequest struct {
	Title       string  `validate:"required,min=1,max=255"`
	Description string  `validate:"required,min=1,max=2000"`
	Rating      float32 `validate:"required,number,gte=0,lte=10"`
	Image       string  `validate:"required,number,min=1,max=255"`
}

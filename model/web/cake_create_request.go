package web

type CakeCreateRequest struct {
	Title       string  `validate:"required,min=1,max=255" json:"title"`
	Description string  `validate:"required,min=1,max=2000" json:"description"`
	Rating      float64 `validate:"required,number,gte=0,lte=10" json:"rating"`
	Image       string  `validate:"required,min=1,max=255" json:"image"`
}

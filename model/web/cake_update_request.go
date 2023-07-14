package web

type CakeUpdateRequest struct {
	Id          int
	Title       string
	Description string
	Rating      float32
	Image       string
}

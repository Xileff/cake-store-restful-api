package web

import "time"

type CakeResponse struct {
	Id          int
	Title       string
	Description string
	Rating      float32
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

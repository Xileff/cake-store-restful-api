package web

type DataResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

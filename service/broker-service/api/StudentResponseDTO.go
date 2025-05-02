package api

type StudentResponseDTO struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

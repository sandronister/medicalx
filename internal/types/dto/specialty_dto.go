package dto

type CreateSpecialtyRequest struct {
	Name string `json:"name"`
}

type UpdateSpecialtyRequest struct {
	Name string `json:"name"`
}

type SpecialtyResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

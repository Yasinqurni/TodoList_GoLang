package dto

type UserRequest struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

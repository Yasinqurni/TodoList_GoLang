package dto

type UserRequest struct {
	Name string `json:"name" binding:"required"`
	Code uint   `json:"code" binding:"required"`
}

type UserLogin struct {
	Code uint `json:"code" binding:"required"`
}

package dto

type TitleRequest struct {
	Title string `json:"title" binding:"required"`
}

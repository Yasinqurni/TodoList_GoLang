package dto

type ActivityRequest struct {
	List string `json:"list" binding:"required"`
	Done bool   `json:"done" binding:"required"`
}

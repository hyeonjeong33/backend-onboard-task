package dto

import (
	"time"
)
type CreateBoardInput struct {
    Title   string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
}

type UpdateBoardInput struct {
    Title   string `json:"title"`
    Content string `json:"content"`
}

type BoardResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Views     uint    `json:"views"`
	CreatedAt time.Time `json:"createdAt"`
}
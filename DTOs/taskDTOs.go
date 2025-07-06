package dtos

import "time"

type TaskReponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type NewTaskRequest struct {
	UserEmail string `json:"email" binding:"required"`
	Title     string `json:"title" binding:"required"`
}

type MarkDoneRequest struct {
	ID   int  `json:"id"`
	Done bool `json:"done"`
}

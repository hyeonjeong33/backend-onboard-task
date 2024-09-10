package models

import "gorm.io/gorm"

type Board struct {
    gorm.Model
    Title   string `json:"title"`
    Content string `json:"content"`
    Views   uint   `json:"views"`
    UserID  uint   `json:"user_id"`
}
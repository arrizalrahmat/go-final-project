package models

import "time"

type Todo struct {
	CommonModels
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
}

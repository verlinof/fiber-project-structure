package example_model

import (
	"time"
)

// ExampleItem represents an example database entity
type ExampleItem struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (ExampleItem) TableName() string {
	return "example_items"
}

package example_service

import (
	"gorm.io/gorm"
)

type ExampleService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) ExampleService {
	return ExampleService{
		db: db,
	}
}

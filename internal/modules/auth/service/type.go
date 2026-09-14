package auth_service

import (
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) AuthService {
	return AuthService{
		db: db,
	}
}

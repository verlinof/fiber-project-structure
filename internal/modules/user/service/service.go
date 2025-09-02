package user_service

import (
	"context"
	"math"

	"github.com/verlinof/fiber-project-structure/db"
	user_model "github.com/verlinof/fiber-project-structure/internal/modules/user/model"
	pkg_success "github.com/verlinof/fiber-project-structure/pkg/success"
	"golang.org/x/crypto/bcrypt"
)

func (s UserService) GetAllUsers(ctx context.Context, page int, perPage int) (*pkg_success.PaginationData, error) {
	var users []user_model.UserResponse
	var totalRows int64

	// Pagination System
	offset := (page - 1) * perPage

	err := db.DB.WithContext(ctx).Table("users").Where("id_role != 1").Count(&totalRows).Limit(perPage).Offset(offset).Find(&users).Error
	if err != nil {
		return &pkg_success.PaginationData{}, err
	}
	totalPage := math.Ceil(float64(totalRows) / float64(perPage))

	response := pkg_success.SuccessPaginationData(users, page, int(totalPage), perPage, int(totalRows))

	return response, nil
}

func (s UserService) GetUserbyPuskesmas(ctx context.Context, id int, page int, perPage int) (*pkg_success.PaginationData, error) {
	var users []user_model.UserResponse
	var totalRows int64

	// Pagination System
	offset := (page - 1) * perPage

	err := db.DB.WithContext(ctx).Table("users").Where("id_role != 1 && id_puskesmas = ?", id).Count(&totalRows).Limit(perPage).Offset(offset).Find(&users).Error
	if err != nil {
		return &pkg_success.PaginationData{}, err
	}

	totalPage := math.Ceil(float64(totalRows) / float64(perPage))
	response := pkg_success.SuccessPaginationData(users, page, int(totalPage), perPage, int(totalRows))

	return response, nil
}

func (s UserService) GetUserbyID(ctx context.Context, id int) (user_model.UserResponse, error) {
	var user user_model.UserResponse

	err := db.DB.WithContext(ctx).Table("users").Where("id = ?", id).First(&user).Error
	if err != nil {
		return user_model.UserResponse{}, err
	}

	return user, nil
}

func (s UserService) CreateUser(ctx context.Context, req user_model.CreateUserRequest) (user_model.UserResponse, error) {
	var userResponse user_model.UserResponse

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return user_model.UserResponse{}, err
	}

	user := user_model.User{
		Username:    req.Username,
		Name:        req.Name,
		Password:    string(hashedPassword),
		IDRole:      req.IDRole,
		IDPuskesmas: req.IDPuskesmas,
		IDPegawai:   req.IDPegawai,
		IDPoli:      req.IDPoli,
		IDPustu:     req.IDPustu,
	}

	err = db.DB.WithContext(ctx).Table("users").Create(&user).Error
	if err != nil {
		return user_model.UserResponse{}, err
	}

	userResponse = user_model.UserResponse{
		ID:          user.ID,
		Username:    user.Username,
		Name:        user.Name,
		IDRole:      user.IDRole,
		IDPuskesmas: user.IDPuskesmas,
		IDPegawai:   user.IDPegawai,
		IDPoli:      user.IDPoli,
		IDPustu:     user.IDPustu,
	}

	return userResponse, nil
}

func (s UserService) ChangePassword(ctx context.Context, id int, req user_model.ChangePasswordRequest) error {
	//Find the data first
	_, err := s.GetUserbyID(ctx, id)
	if err != nil {
		return err
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = db.DB.WithContext(ctx).Table("users").Where("id = ?", id).Update("password", string(hashedPassword)).Error
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) UpdateUser(ctx context.Context, id int, req user_model.UpdateUserRequest) (user_model.UserResponse, error) {
	var userResponse user_model.UserResponse

	//Find the data first
	_, err := s.GetUserbyID(ctx, id)
	if err != nil {
		return user_model.UserResponse{}, err
	}

	err = db.DB.WithContext(ctx).Table("users").Where("id = ?", id).Updates(req).Error
	if err != nil {
		return user_model.UserResponse{}, err
	}

	userResponse = user_model.UserResponse{
		ID:          id,
		Username:    req.Username,
		Name:        req.Name,
		IDRole:      req.IDRole,
		IDPuskesmas: req.IDPuskesmas,
		IDPegawai:   req.IDPegawai,
		IDPoli:      req.IDPoli,
		IDPustu:     req.IDPustu,
	}

	return userResponse, nil
}

func (s UserService) DeleteUser(ctx context.Context, id int) error {
	//Find the data first
	_, err := s.GetUserbyID(ctx, id)
	if err != nil {
		return err
	}

	err = db.DB.WithContext(ctx).Table("users").Where("id = ?", id).Delete(&user_model.User{}).Error
	if err != nil {
		return err
	}

	return nil
}

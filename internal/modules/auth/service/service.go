package auth_service

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/verlinof/fiber-project-structure/db"
	auth_model "github.com/verlinof/fiber-project-structure/internal/modules/auth/model"
	user_model "github.com/verlinof/fiber-project-structure/internal/modules/user/model"
	pkg_jwt "github.com/verlinof/fiber-project-structure/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

func (authService *AuthService) Login(ctx context.Context, loginRequest auth_model.LoginRequest) (auth_model.LoginResponse, error) {
	var authModel user_model.User
	var loginResponse auth_model.LoginResponse
	var userResponse user_model.UserResponse

	err := db.DB.WithContext(ctx).Table("users").Where("username = ?", loginRequest.Username).First(&authModel).Error
	if err != nil {
		return auth_model.LoginResponse{}, err
	}

	// Compare Password
	err = bcrypt.CompareHashAndPassword([]byte(authModel.Password), []byte(loginRequest.Password))
	if err != nil {
		return auth_model.LoginResponse{}, err
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"id_user": authModel.ID,
		"id_role": authModel.IDRole,
	}

	token, err := pkg_jwt.GenerateJWT(&claims)
	if err != nil {
		return auth_model.LoginResponse{}, err
	}

	userResponse = user_model.UserResponse{
		ID:          authModel.ID,
		Username:    authModel.Username,
		Name:        authModel.Name,
		IDRole:      authModel.IDRole,
		IDPuskesmas: authModel.IDPuskesmas,
		IDPegawai:   authModel.IDPegawai,
		IDPoli:      authModel.IDPoli,
		IDPustu:     authModel.IDPustu,
	}

	loginResponse = auth_model.LoginResponse{
		Jwt:  token,
		User: &userResponse,
	}

	return loginResponse, nil
}

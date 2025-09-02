package auth_model

import user_model "github.com/verlinof/fiber-project-structure/internal/modules/user/model"

type LoginResponse struct {
	Jwt  string                   `json:"jwt"`
	User *user_model.UserResponse `json:"user,omitempty"`
}

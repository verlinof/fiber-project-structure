package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/verlinof/fiber-project-structure/db"
	user_model "github.com/verlinof/fiber-project-structure/internal/modules/user/model"
	pkg_error "github.com/verlinof/fiber-project-structure/pkg/error"
)

// INI SEMENTARA, KALAU MODULNYA DAH ADA BAKAL DIILANGIN
type permission struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type roleHasPermissions struct {
	RoleId       int `json:"role_id"`
	PermissionId int `json:"permission_id"`
}

// RoleMiddleware enforces permission-based access control
func RoleMiddleware(tag string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var permission permission
		var roleHasPermission roleHasPermissions
		var userModel user_model.User

		user := c.Locals("token").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		// Get Data From Token
		IDUser := claims["id_user"].(float64)
		IDRole := claims["id_role"].(float64)

		fmt.Println(IDUser, IDRole)

		// Find Permission by Tag
		err := db.DB.Table("permissions").Where("name = ?", tag).First(&permission).Error
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(pkg_error.NewForbidden(fmt.Errorf("access denied")))
		}

		// Check Role Has Permission
		err = db.DB.Table("role_has_permissions").Where("role_id = ? AND permission_id = ?", IDRole, permission.ID).First(&roleHasPermission).Error
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(pkg_error.NewForbidden(fmt.Errorf("access denied")))
		}

		// Set Current User
		err = db.DB.Table("users").Where("id = ?", IDUser).First(&userModel).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(pkg_error.NewInternalServerError(err))
		}
		// Set User Locals
		c.Locals("user", userModel)

		return c.Next()
	}
}

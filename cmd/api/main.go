package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/verlinof/fiber-project-structure/configs/app_config"
	"github.com/verlinof/fiber-project-structure/configs/db_config"
	"github.com/verlinof/fiber-project-structure/configs/redis_config"
	"github.com/verlinof/fiber-project-structure/db"
	"github.com/verlinof/fiber-project-structure/internal/route"
)

func main() {
	//Init Global Config
	app_config.Config = app_config.LoadConfig()
	db_config.Config = db_config.LoadConfig()
	redis_config.Config = redis_config.LoadConfig()

	//Connect Database
	db.ConnectDatabase()

	// Init Fiber Engine
	router := fiber.New()

	// Init Route
	route.InitRoute(router)

	//Run Server
	router.Listen(":" + app_config.Config.AppPort)

	// //Init GIN ENGINE
	// gin.SetMode(app_config.Config.GinMode)
	// router := gin.Default()

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour, // How long preflight requests can be cached
	// }))

	// //Init Router
	// route.InitRoute(router)

	// //Run Server
	// router.Run(":" + app_config.Config.AppPort)
}

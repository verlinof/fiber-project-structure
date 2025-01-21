package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/verlinof/fiber-project-structure/configs/app_config"
	"github.com/verlinof/fiber-project-structure/configs/db_config"
	"github.com/verlinof/fiber-project-structure/configs/redis_config"
	"github.com/verlinof/fiber-project-structure/db"
	"github.com/verlinof/fiber-project-structure/internal/routes"
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
	routes.InitRoute(router)

	//Run Server
	router.Listen(":" + app_config.Config.AppPort)
}

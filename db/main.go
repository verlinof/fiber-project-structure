package db

import (
	"github.com/verlinof/fiber-project-structure/configs/db_config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	dsnMysql := db_config.Config.DbUser + ":" + db_config.Config.DbPassword + "@tcp(" + db_config.Config.Host + ":" + db_config.Config.Port + ")/" + db_config.Config.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsnMysql), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

package db

import (
	"fmt"
	"log"
	"time"

	"github.com/verlinof/fiber-project-structure/configs/db_config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Init() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		db_config.Config.Host,
		db_config.Config.DbUser,
		db_config.Config.DbPassword,
		db_config.Config.DbName,
		db_config.Config.Port,
	)

	// Menambahkan Konfigurasi Logger GORM (saran perbaikan)
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			ParameterizedQueries:      true,
			SlowThreshold:             time.Second, // Ambang batas query lambat
			LogLevel:                  logger.Warn, // Level log (Info, Warn, Error)
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Buka koneksi dan assign ke variabel global DB
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		// Jika koneksi database gagal saat startup, aplikasi tidak bisa berjalan.
		// Jadi, kita gunakan panic di sini agar aplikasi berhenti.
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	log.Println("Koneksi database berhasil dibuat.")

	// Konfigurasi Connection Pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Gagal mendapatkan instance sql.DB dari GORM: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

// GetDB adalah "getter" yang Anda tanyakan.
func GetDB() *gorm.DB {
	return db
}

package pkg_cron

import (
	"gorm.io/gorm"
)

// RegisterJobs adalah tempat untuk mendaftarkan semua cron job aplikasi.
func RegisterJobs(s *Scheduler, db *gorm.DB) {
	// Contoh pendaftaran cron job:
	// s.AddJob("0 1 * * *", func() {
	//     // Logic cron job di sini
	// })
}

package pkg_cron

import (
	"log"

	"github.com/robfig/cron/v3"
)

// Scheduler mengelola semua tugas terjadwal (cron jobs).
type Scheduler struct {
	cron *cron.Cron
}

func NewScheduler() *Scheduler {
	c := cron.New(cron.WithChain(cron.Recover(cron.DefaultLogger)))
	return &Scheduler{cron: c}
}

// AddJob mendaftarkan sebuah fungsi untuk dijalankan sesuai jadwal (spec).
// spec: Format cron (e.g., "0 1 * * *" untuk jam 1 pagi setiap hari).
// cmd: Fungsi yang akan dijalankan (tanpa argumen).
func (s *Scheduler) AddJob(spec string, cmd func()) (cron.EntryID, error) {
	id, err := s.cron.AddFunc(spec, cmd)
	if err != nil {
		log.Printf("[CRON-ERROR] Failed to add job with spec '%s': %v", spec, err)
		return 0, err
	}
	log.Printf("[CRON] Job registered with spec '%s', ID: %d", spec, id)
	return id, nil
}

// Start menjalankan scheduler di background.
func (s *Scheduler) Start() {
	log.Println("[CRON] Starting scheduler...")
	s.cron.Start()
	log.Println("[CRON] Scheduler started successfully.")
}

// Stop menghentikan scheduler dan menunggu job yang sedang berjalan untuk selesai.
// Ini penting untuk graceful shutdown.
func (s *Scheduler) Stop() {
	log.Println("[CRON] Stopping scheduler...")
	ctx := s.cron.Stop()
	<-ctx.Done() // Menunggu semua job yang sedang berjalan selesai.
	log.Println("[CRON] Scheduler stopped gracefully.")
}

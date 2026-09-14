package pkg_email

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"sync"
	"time"

	"gopkg.in/gomail.v2"
)

// SmtpConfig tetap sama
type SmtpConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Sender   string
}

// EmailOTPData sekarang menampung nama dan kode OTP
type EmailOTPData struct {
	Name string
	Code string
}

func NewSmtpConfig() SmtpConfig {
	return SmtpConfig{
		Host:     os.Getenv("MAIL_HOST"),
		Port:     os.Getenv("MAIL_PORT"),
		User:     os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
		Sender:   os.Getenv("MAIL_FROM_ADDRESS"),
	}
}

// EmailJob adalah struktur untuk job email
type EmailJob struct {
	RecipientName  string
	RecipientEmail string
	Code           string
	Timestamp      time.Time
}

// EmailWorker adalah worker pool untuk mengirim email
type EmailWorker struct {
	jobs       chan EmailJob
	workers    int
	wg         sync.WaitGroup
	maxRetries int
}

var (
	globalEmailWorker *EmailWorker
	once              sync.Once
)

// NewEmailWorker membuat worker pool baru
func NewEmailWorker(workers int, queueSize int, maxRetries int) *EmailWorker {
	return &EmailWorker{
		jobs:       make(chan EmailJob, queueSize),
		workers:    workers,
		maxRetries: maxRetries,
	}
}

// InitEmailWorker inisialisasi global email worker (dipanggil saat aplikasi start)
func InitEmailWorker(workers, queueSize, maxRetries int) {
	once.Do(func() {
		globalEmailWorker = NewEmailWorker(workers, queueSize, maxRetries)
		globalEmailWorker.Start()
	})
}

// Start menjalankan worker pool
func (w *EmailWorker) Start() {
	for i := 0; i < w.workers; i++ {
		w.wg.Add(1)
		go w.worker(i)
	}
}

// worker adalah goroutine yang memproses email job
func (w *EmailWorker) worker(id int) {
	defer w.wg.Done()

	for job := range w.jobs {
		log.Printf("[Worker-%d] Processing email to %s (queued at %s)",
			id, job.RecipientEmail, job.Timestamp.Format("15:04:05"))

		// Retry mechanism dengan exponential backoff
		var err error
		for attempt := 1; attempt <= w.maxRetries; attempt++ {
			err = sendEmailSync(job.RecipientName, job.RecipientEmail, job.Code)

			if err == nil {
				log.Printf("[Worker-%d] Email successfully sent to %s", id, job.RecipientEmail)
				break
			}

			log.Printf("[Worker-%d] Attempt %d/%d failed: %v",
				id, attempt, w.maxRetries, err)

			// Exponential backoff sebelum retry
			if attempt < w.maxRetries {
				backoffDuration := time.Duration(attempt*attempt) * time.Second
				log.Printf("[Worker-%d] Retrying in %v...", id, backoffDuration)
				time.Sleep(backoffDuration)
			}
		}

		if err != nil {
			log.Printf("[Worker-%d] ✗✗ FAILED to send email to %s after %d attempts: %v",
				id, job.RecipientEmail, w.maxRetries, err)
			// TODO: Bisa simpan ke database untuk manual retry atau kirim alert ke admin
		}
	}

	log.Printf("[Worker-%d] Shutting down", id)
}

// sendEmailSync adalah fungsi internal untuk mengirim email (synchronous)
func sendEmailSync(recipientName, recipientEmail, code string) error {
	config := NewSmtpConfig()

	// Parse template
	templatePath := "./pkg/email/email_otp.html"
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("gagal membaca template email: %w", err)
	}

	data := EmailOTPData{
		Name: recipientName,
		Code: code,
	}

	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return fmt.Errorf("gagal mengeksekusi template email: %w", err)
	}

	// Setup message
	m := gomail.NewMessage()
	m.SetHeader("From", config.Sender)
	m.SetHeader("To", recipientEmail)
	m.SetHeader("Subject", "Kode Verifikasi Akun Anda")
	m.SetBody("text/html", body.String())

	port, err := strconv.Atoi(config.Port)
	if err != nil {
		return fmt.Errorf("port tidak valid: %w", err)
	}

	// Setup dialer
	d := gomail.NewDialer(config.Host, port, config.User, config.Password)

	if port == 465 {
		d.SSL = true
	}

	// TLS Config (untuk production, ubah InsecureSkipVerify ke false)
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true, // GANTI jadi false di production!
		ServerName:         config.Host,
	}

	// Timeout mechanism
	done := make(chan error, 1)
	go func() {
		done <- d.DialAndSend(m)
	}()

	select {
	case err := <-done:
		return err
	case <-time.After(30 * time.Second):
		return fmt.Errorf("timeout: koneksi SMTP lebih dari 30 detik")
	}
}

// SendEmailAsync menambahkan email ke queue (non-blocking)
func SendEmailAsync(recipientName, recipientEmail, code string) error {
	if globalEmailWorker == nil {
		return fmt.Errorf("email worker belum diinisialisasi, panggil InitEmailWorker() terlebih dahulu")
	}

	job := EmailJob{
		RecipientName:  recipientName,
		RecipientEmail: recipientEmail,
		Code:           code,
		Timestamp:      time.Now(),
	}

	// Non-blocking send dengan timeout
	select {
	case globalEmailWorker.jobs <- job:
		log.Printf("✓ Email job queued for %s", recipientEmail)
		return nil
	case <-time.After(2 * time.Second):
		return fmt.Errorf("email queue penuh, tidak bisa menambahkan job")
	}
}

// SendEmailSync untuk backward compatibility (kirim langsung tanpa queue)
func SendEmailSync(recipientName, recipientEmail, code string) error {
	return sendEmailSync(recipientName, recipientEmail, code)
}

// ShutdownEmailWorker untuk graceful shutdown
func ShutdownEmailWorker() {
	if globalEmailWorker != nil {
		log.Println("Shutting down email worker...")
		close(globalEmailWorker.jobs) // Stop menerima job baru
		globalEmailWorker.wg.Wait()   // Tunggu semua job selesai
		log.Println("✓ Email worker stopped gracefully")
	}
}

// GetQueueSize mengembalikan jumlah email yang sedang di queue
func GetQueueSize() int {
	if globalEmailWorker == nil {
		return 0
	}
	return len(globalEmailWorker.jobs)
}

func SendVerificationCode(recipientName, recipientEmail, code string) error {
	config := NewSmtpConfig()
	fmt.Printf("=== DEBUG EMAIL ===\n")
	fmt.Printf("Host: %s\n", config.Host)
	fmt.Printf("Port: %s\n", config.Port)
	fmt.Printf("User: %s\n", config.User)
	fmt.Printf("To: %s\n", recipientEmail)

	return SendEmailAsync(recipientName, recipientEmail, code)
}

func SendVerificationCodeTLS(recipientName, recipientEmail, code string) error {
	config := NewSmtpConfig()
	templatePath := "./pkg/email/email_otp.html"

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("gagal membaca template email: %w", err)
	}

	data := EmailOTPData{
		Name: recipientName,
		Code: code,
	}

	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return fmt.Errorf("gagal mengeksekusi template email: %w", err)
	}

	// Buat pesan email lengkap (termasuk header)
	headers := make(map[string]string)
	headers["From"] = config.Sender
	headers["To"] = recipientEmail
	headers["Subject"] = "Kode Verifikasi Akun Anda"
	headers["MIME-version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"UTF-8\""

	var msg string
	for k, v := range headers {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n" + body.String()

	// 1. Buat alamat server
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)

	// 2. Konfigurasi otentikasi
	auth := smtp.PlainAuth("", config.User, config.Password, config.Host)

	// 3. Buat konfigurasi TLS
	// InsecureSkipVerify: true seringkali diperlukan saat development dengan sertifikat self-signed.
	// Untuk produksi, sebaiknya diatur ke false dan pastikan server memiliki sertifikat yang valid.
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true, // Hati-hati di produksi!
		ServerName:         config.Host,
	}

	// 4. Buat koneksi TCP yang langsung dienkripsi dengan TLS
	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return fmt.Errorf("gagal membuat koneksi TLS: %w", err)
	}
	defer conn.Close()

	// 5. Buat klien SMTP baru dari koneksi yang sudah terenkripsi
	client, err := smtp.NewClient(conn, config.Host)
	if err != nil {
		return fmt.Errorf("gagal membuat klien SMTP: %w", err)
	}
	defer client.Quit()

	// 6. Lakukan otentikasi
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("gagal otentikasi SMTP: %w", err)
	}

	// 7. Atur pengirim (MAIL FROM)
	if err = client.Mail(config.Sender); err != nil {
		return fmt.Errorf("gagal mengatur pengirim: %w", err)
	}

	// 8. Atur penerima (RCPT TO)
	if err = client.Rcpt(recipientEmail); err != nil {
		return fmt.Errorf("gagal mengatur penerima: %w", err)
	}

	// 9. Dapatkan writer untuk menulis isi email
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("gagal membuka stream data: %w", err)
	}

	// 10. Tulis pesan dan tutup writer
	_, err = w.Write([]byte(msg))
	if err != nil {
		return fmt.Errorf("gagal menulis isi email: %w", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("gagal menutup stream data: %w", err)
	}

	// ---- AKHIR BAGIAN YANG BERUBAH ----

	return nil
}

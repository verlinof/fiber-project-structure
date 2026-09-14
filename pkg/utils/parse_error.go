package pkg_utils

import (
	"errors"
	"log"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func ParseError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("data not found")
	}

	if pqErr, ok := err.(*pq.Error); ok {
		log.Printf("Error Database Terdeteksi: Kode %s", pqErr.Code) // Log untuk developer

		switch pqErr.Code {
		case "23505": // Unique Violation
			// Kita bisa parsing 'pqErr.Detail' untuk tahu field mana,
			// tapi untuk simpelnya, kita beri pesan general
			return errors.New("data exist")

		case "23503": // Foreign Key Violation
			return errors.New("invalid foreign key reference")

		case "23502": // Not Null Violation
			return errors.New("missing required field")

			// Anda bisa tambahkan kode error Postgres lainnya di sini
			// https://www.postgresql.org/docs/current/errcodes-appendix.html
		}
	}

	// 3. Fallback untuk error umum
	// Kita log error aslinya untuk debugging
	log.Println("Error tidak terduga:", err.Error())
	return errors.New("internal server error")
}

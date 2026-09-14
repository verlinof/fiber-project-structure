package pkg_utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
)

// GenerateOTPCode menghasilkan kode numerik acak dengan panjang yang ditentukan.
// Untuk keamanan, gunakan panjang 6 digit.
func GenerateOTPCode() (string, error) {
	length := 6
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		return "", err
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b), nil
}

// Order Number Generator
// Material Order Number Format: ORD-MAT-{timestamp}
func GenerateOrderNumber() string {
	now := time.Now()
	shortUUID := uuid.New().String()[:8]

	return fmt.Sprintf("ORD-MAT-%s-%s",
		now.Format("20060102"),
		shortUUID)
}

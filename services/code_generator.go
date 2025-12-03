package services

import (
	"fmt"
	"time"
)

// generateCode membuat kode unik berbasis prefix dan timestamp nano.
// Disimpan di level package agar bisa digunakan lintas service.
func generateCode(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, time.Now().UnixNano())
}




package models

import "time"

type KehadiranPeminjam struct {
	KodeKehadiran   string               `json:"kode_kehadiran"`
	KodePeminjaman  string               `json:"kode_peminjaman"`
	Peminjaman      *Peminjaman          `json:"peminjaman,omitempty"`
	StatusKehadiran KehadiranStatusEnum  `json:"status_kehadiran"`
	WaktuKonfirmasi time.Time            `json:"waktu_konfirmasi"`
	Keterangan      string               `json:"keterangan"`
	DiverifikasiOleh *string             `json:"diverifikasi_oleh"`
	Verifier        *User                `json:"verifier,omitempty"`
	CreatedAt       time.Time            `json:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at"`
}

type CreateKehadiranRequest struct {
	KodePeminjaman  string              `json:"kode_peminjaman"`
	StatusKehadiran KehadiranStatusEnum `json:"status_kehadiran"`
	Keterangan      string              `json:"keterangan"`
}

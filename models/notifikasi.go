package models

import "time"

type Notifikasi struct {
	KodeNotifikasi   string               `json:"kode_notifikasi"`
	KodeUser         string               `json:"kode_user"`
	KodePeminjaman   *string              `json:"kode_peminjaman"`
	JenisNotifikasi  NotifikasiJenisEnum  `json:"jenis_notifikasi"`
	Pesan            string               `json:"pesan"`
	Status           NotifikasiStatusEnum `json:"status"`
	CreatedAt        time.Time            `json:"created_at"`
	UpdatedAt        *time.Time           `json:"updated_at"`
	Peminjaman       *Peminjaman          `json:"peminjaman,omitempty"`
}


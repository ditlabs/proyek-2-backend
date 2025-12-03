package models

import "time"

type LogAktivitas struct {
	KodeLog        string       `json:"kode_log"`
	KodeUser       *string      `json:"kode_user"`
	User           *User        `json:"user,omitempty"`
	KodePeminjaman *string      `json:"kode_peminjaman"`
	Peminjaman     *Peminjaman  `json:"peminjaman,omitempty"`
	Aksi           string       `json:"aksi"`
	Keterangan     string       `json:"keterangan"`
	Waktu          time.Time    `json:"waktu"`
	UpdatedAt      *time.Time   `json:"updated_at"`
}


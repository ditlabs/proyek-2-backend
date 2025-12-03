package models

import "time"

type Kegiatan struct {
	KodeKegiatan   string     `json:"kode_kegiatan"`
	NamaKegiatan   string     `json:"nama_kegiatan"`
	Deskripsi      string     `json:"deskripsi"`
	TanggalMulai   time.Time  `json:"tanggal_mulai"`
	TanggalSelesai time.Time  `json:"tanggal_selesai"`
	OrganisasiKode string     `json:"organisasi_kode"`
	Organisasi     *Organisasi `json:"organisasi,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}



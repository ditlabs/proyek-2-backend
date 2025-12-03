package models

import "time"

type Organisasi struct {
	KodeOrganisasi  string              `json:"kode_organisasi"`
	Nama            string              `json:"nama"`
	JenisOrganisasi JenisOrganisasiEnum `json:"jenis_organisasi"`
	Kontak          string              `json:"kontak"`
	CreatedAt       time.Time           `json:"created_at"`
}


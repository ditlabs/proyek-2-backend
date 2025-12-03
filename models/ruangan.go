package models

type Ruangan struct {
	KodeRuangan string `json:"kode_ruangan"`
	NamaRuangan string `json:"nama_ruangan"`
	Lokasi      string `json:"lokasi"`
	Kapasitas   int    `json:"kapasitas"`
	Deskripsi   string `json:"deskripsi"`
}

type CreateRuanganRequest struct {
	KodeRuangan string `json:"kode_ruangan"`
	NamaRuangan string `json:"nama_ruangan"`
	Lokasi      string `json:"lokasi"`
	Kapasitas   int    `json:"kapasitas"`
	Deskripsi   string `json:"deskripsi"`
}

type UpdateRuanganRequest struct {
	NamaRuangan string `json:"nama_ruangan"`
	Lokasi      string `json:"lokasi"`
	Kapasitas   int    `json:"kapasitas"`
	Deskripsi   string `json:"deskripsi"`
}


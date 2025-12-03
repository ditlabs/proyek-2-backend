package handlers

import (
	"backend-sarpras/models"
	"backend-sarpras/repositories"
	"encoding/json"
	"net/http"
	"strings"
)

type BarangHandler struct {
	BarangRepo *repositories.BarangRepository
}

func NewBarangHandler(barangRepo *repositories.BarangRepository) *BarangHandler {
	return &BarangHandler{BarangRepo: barangRepo}
}

func (h *BarangHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	barangs, err := h.BarangRepo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(barangs)
}

func (h *BarangHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	kode := strings.TrimPrefix(r.URL.Path, "/api/barang/")
	if kode == "" {
		http.Error(w, "Invalid kode_barang", http.StatusBadRequest)
		return
	}

	barang, err := h.BarangRepo.GetByID(kode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if barang == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(barang)
}

func (h *BarangHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.CreateBarangRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	barang := &models.Barang{
		KodeBarang:  req.KodeBarang,
		NamaBarang:  req.NamaBarang,
		Deskripsi:   req.Deskripsi,
		JumlahTotal: req.JumlahTotal,
		RuanganKode: req.RuanganKode,
	}

	if err := h.BarangRepo.Create(barang); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(barang)
}

func (h *BarangHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	kode := strings.TrimPrefix(r.URL.Path, "/api/barang/")
	if kode == "" {
		http.Error(w, "Invalid kode_barang", http.StatusBadRequest)
		return
	}

	var req models.UpdateBarangRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	barang := &models.Barang{
		KodeBarang:  kode,
		NamaBarang:  req.NamaBarang,
		Deskripsi:   req.Deskripsi,
		JumlahTotal: req.JumlahTotal,
		RuanganKode: req.RuanganKode,
	}

	if err := h.BarangRepo.Update(barang); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(barang)
}

func (h *BarangHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	kode := strings.TrimPrefix(r.URL.Path, "/api/barang/")
	if kode == "" {
		http.Error(w, "Invalid kode_barang", http.StatusBadRequest)
		return
	}

	if err := h.BarangRepo.Delete(kode); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}


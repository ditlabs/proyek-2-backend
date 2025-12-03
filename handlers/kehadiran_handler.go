package handlers

import (
	"backend-sarpras/middleware"
	"backend-sarpras/models"
	"backend-sarpras/repositories"
	"backend-sarpras/services"
	"encoding/json"
	"net/http"
)

type KehadiranHandler struct {
	KehadiranService *services.KehadiranService
	KehadiranRepo    *repositories.KehadiranRepository
	PeminjamanRepo   *repositories.PeminjamanRepository
	RuanganRepo      *repositories.RuanganRepository
	UserRepo         *repositories.UserRepository
}

func NewKehadiranHandler(
	kehadiranService *services.KehadiranService,
	kehadiranRepo *repositories.KehadiranRepository,
	peminjamanRepo *repositories.PeminjamanRepository,
	ruanganRepo *repositories.RuanganRepository,
	userRepo *repositories.UserRepository,
) *KehadiranHandler {
	return &KehadiranHandler{
		KehadiranService: kehadiranService,
		KehadiranRepo:    kehadiranRepo,
		PeminjamanRepo:   peminjamanRepo,
		RuanganRepo:      ruanganRepo,
		UserRepo:         userRepo,
	}
}

func (h *KehadiranHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user := middleware.GetUserFromContext(r)
	if user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req models.CreateKehadiranRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.KehadiranService.CreateKehadiran(&req, user.KodeUser)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Kehadiran berhasil dicatat"})
}

func (h *KehadiranHandler) GetByPeminjamanID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	peminjamanIDStr := r.URL.Query().Get("peminjaman_id")
	if peminjamanIDStr == "" {
		http.Error(w, "peminjaman_id required", http.StatusBadRequest)
		return
	}

	// TODO: parse peminjamanID dari query string
	// Untuk sekarang, return empty atau implementasi sesuai kebutuhan
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]models.KehadiranPeminjam{})
}

func (h *KehadiranHandler) GetRiwayatBySecurity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	user := middleware.GetUserFromContext(r)
	if user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Saat ini repository GetBySecurityID belum mengembalikan data terfilter (stub),
	// jadi kita cukup mengembalikan hasil mentah tanpa enrich tambahan.
	kehadiranList, err := h.KehadiranRepo.GetBySecurityID(0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kehadiranList)
}

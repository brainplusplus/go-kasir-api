package handler

import (
	"encoding/json"
	"kasir-api/internal/adapter/inbound/http/dto"
	"kasir-api/internal/port/inbound"
	"kasir-api/internal/utils/swagger"
	"net/http"
	"strconv"
	"strings"
)

type ProdukHandler struct {
	service inbound.ProdukService
	swag    *swagger.Generator
}

func NewProdukHandler(service inbound.ProdukService, swag *swagger.Generator) *ProdukHandler {
	return &ProdukHandler{service: service, swag: swag}
}

func (h *ProdukHandler) RegisterRoutes(mux *http.ServeMux) {
	// Register to Swagger
	h.swag.Register("GET", "/api/produk", "Get All Products", nil, []dto.ProdukResponse{})
	h.swag.Register("POST", "/api/produk", "Create New Product", dto.CreateProdukRequest{}, dto.ProdukResponse{})
	h.swag.Register("GET", "/api/produk/{id}", "Get Product By ID", nil, dto.ProdukResponse{})
	h.swag.Register("PUT", "/api/produk/{id}", "Update Product", dto.UpdateProdukRequest{}, dto.ProdukResponse{})
	h.swag.Register("DELETE", "/api/produk/{id}", "Delete Product", nil, map[string]string{"message": "sukses"})

	// Register to Mux
	mux.HandleFunc("/api/produk/", h.HandleProduk)
	mux.HandleFunc("/api/produk", h.HandleProduk)
}

func (h *ProdukHandler) HandleProduk(w http.ResponseWriter, r *http.Request) {
	// /api/produk/
	if strings.HasPrefix(r.URL.Path, "/api/produk/") && len(r.URL.Path) > len("/api/produk/") {
		idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
			return
		}

		if r.Method == "GET" {
			h.getByID(w, id)
		} else if r.Method == "PUT" {
			h.update(w, r, id)
		} else if r.Method == "DELETE" {
			h.delete(w, id)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	// /api/produk
	if r.URL.Path == "/api/produk" || r.URL.Path == "/api/produk/" {
		if r.Method == "GET" {
			h.getAll(w)
		} else if r.Method == "POST" {
			h.create(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		return
	}

	http.NotFound(w, r)
}

func (h *ProdukHandler) getAll(w http.ResponseWriter) {
	products, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responses := make([]dto.ProdukResponse, 0)
	for _, p := range products {
		responses = append(responses, dto.FromDomainProduk(p))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

func (h *ProdukHandler) getByID(w http.ResponseWriter, id int) {
	p, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "Produk belum ada", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.FromDomainProduk(*p))
}

func (h *ProdukHandler) create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProdukRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	created, err := h.service.Create(req.ToDomain())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto.FromDomainProduk(created))
}

func (h *ProdukHandler) update(w http.ResponseWriter, r *http.Request, id int) {
	var req dto.UpdateProdukRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	updated, err := h.service.Update(id, req.ToDomain())
	if err != nil {
		http.Error(w, "Produk belum ada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.FromDomainProduk(*updated))
}

func (h *ProdukHandler) delete(w http.ResponseWriter, id int) {
	err := h.service.Delete(id)
	if err != nil {
		http.Error(w, "Produk belum ada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "sukses delete",
	})
}

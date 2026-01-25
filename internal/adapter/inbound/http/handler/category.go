package handler

import (
	"encoding/json"
	"kasir-api/internal/model"
	"kasir-api/internal/port/inbound"
	"kasir-api/internal/utils/swagger"
	"net/http"
	"strconv"
	"strings"
)

type CategoryHandler struct {
	service inbound.CategoryService
	swag    *swagger.Generator
}

func NewCategoryHandler(service inbound.CategoryService, swag *swagger.Generator) *CategoryHandler {
	return &CategoryHandler{service: service, swag: swag}
}

func (h *CategoryHandler) RegisterRoutes(mux *http.ServeMux) {
	// Register to Swagger
	h.swag.Register("GET", "/api/categories", "Get All Categories", nil, []model.Category{})
	h.swag.Register("POST", "/api/categories", "Create New Category", model.Category{}, model.Category{})
	h.swag.Register("GET", "/api/categories/{id}", "Get Category By ID", nil, model.Category{})
	h.swag.Register("PUT", "/api/categories/{id}", "Update Category", model.Category{}, model.Category{})
	h.swag.Register("DELETE", "/api/categories/{id}", "Delete Category", nil, map[string]string{"message": "sukses"})

	// Register to Mux
	mux.HandleFunc("/api/categories/", h.HandleCategory)
	mux.HandleFunc("/api/categories", h.HandleCategory)
}

func (h *CategoryHandler) HandleCategory(w http.ResponseWriter, r *http.Request) {
	// /api/categories/
	if strings.HasPrefix(r.URL.Path, "/api/categories/") && len(r.URL.Path) > len("/api/categories/") {
		idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid Category ID", http.StatusBadRequest)
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

	// /api/categories
	if r.URL.Path == "/api/categories" || r.URL.Path == "/api/categories/" {
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

func (h *CategoryHandler) getAll(w http.ResponseWriter) {
	categories, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (h *CategoryHandler) getByID(w http.ResponseWriter, id int) {
	category, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

func (h *CategoryHandler) create(w http.ResponseWriter, r *http.Request) {
	var c model.Category
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	created, err := h.service.Create(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *CategoryHandler) update(w http.ResponseWriter, r *http.Request, id int) {
	var c model.Category
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	updated, err := h.service.Update(id, c)
	if err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func (h *CategoryHandler) delete(w http.ResponseWriter, id int) {
	err := h.service.Delete(id)
	if err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "sukses delete",
	})
}

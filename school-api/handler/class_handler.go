package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"school-api/models"
	"school-api/service"
	"github.com/gorilla/mux"
)

type ClassHandler struct {
	service service.ClassService
}

func NewClassHandler(service service.ClassService) *ClassHandler {
	return &ClassHandler{service: service}
}

// @Summary Create a new class
// @Description Create a new class with the provided details
// @Tags classes
// @Accept json
// @Produce json
// @Param class body models.Class true "Class object to create"
// @Success 201 {object} models.Class
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /classes [post]
func (h *ClassHandler) CreateClass(w http.ResponseWriter, r *http.Request) {
	var class models.Class
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateClass(&class); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(class)
}

// @Summary Get all classes
// @Description Get a list of all classes
// @Tags classes
// @Produce json
// @Success 200 {array} models.Class
// @Failure 500 {string} string "Internal server error"
// @Router /classes [get]
func (h *ClassHandler) GetAllClasses(w http.ResponseWriter, r *http.Request) {
	classes, err := h.service.GetAllClasses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classes)
}

// @Summary Get a class by ID
// @Description Get a specific class by its ID
// @Tags classes
// @Produce json
// @Param id path int true "Class ID"
// @Success 200 {object} models.Class
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Class not found"
// @Router /classes/{id} [get]
func (h *ClassHandler) GetClassByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	class, err := h.service.GetClassByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(class)
} 
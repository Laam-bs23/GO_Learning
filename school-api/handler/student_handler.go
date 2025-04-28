package handler

import (
	"encoding/json"
	"net/http"
	"school-api/models"
	"school-api/service"
	"strconv"

	"github.com/gorilla/mux"
)

type StudentHandler interface {
	CreateStudent(w http.ResponseWriter, r *http.Request)
	GetAllStudents(w http.ResponseWriter, r *http.Request)
	GetStudentByID(w http.ResponseWriter, r *http.Request)
	UpdateStudent(w http.ResponseWriter, r *http.Request)
	DeleteStudent(w http.ResponseWriter, r *http.Request)
}

type studentHandler struct {
	studentService service.StudentService
}

func NewStudentHandler(studentService service.StudentService) StudentHandler {
	return &studentHandler{
		studentService: studentService,
	}
}

// @Summary Create a new student
// @Description Create a new student with the provided details
// @Tags students
// @Accept json
// @Produce json
// @Param student body models.Student true "Student object to create"
// @Success 201 {object} models.Student
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /students [post]
func (h *studentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.studentService.CreateStudent(&student); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

// @Summary Get all students
// @Description Get a list of all students
// @Tags students
// @Produce json
// @Success 200 {array} models.Student
// @Failure 500 {string} string "Internal server error"
// @Router /students [get]
func (h *studentHandler) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := h.studentService.GetAllStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// @Summary Get a student by ID
// @Description Get a specific student by its ID
// @Tags students
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {object} models.Student
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Student not found"
// @Router /students/{id} [get]
func (h *studentHandler) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	student, err := h.studentService.GetStudentByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// @Summary Update a student
// @Description Update an existing student with the provided details
// @Tags students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Param student body models.Student true "Student object to update"
// @Success 200 {object} models.Student
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Internal server error"
// @Router /students/{id} [put]
func (h *studentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	student.ID = uint(id)
	if err := h.studentService.UpdateStudent(&student); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// @Summary Delete a student
// @Description Delete a specific student by its ID
// @Tags students
// @Param id path int true "Student ID"
// @Success 204 "No Content"
// @Failure 400 {string} string "Invalid ID"
// @Failure 500 {string} string "Internal server error"
// @Router /students/{id} [delete]
func (h *studentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	if err := h.studentService.DeleteStudent(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
} 
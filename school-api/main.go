package main

import (
	"log"
	"net/http"
	"school-api/docs"
	"school-api/handler"
	"school-api/models"
	"school-api/repository"
	"school-api/service"
	"time"
	"os/exec"
	"runtime"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// @title School API
// @version 1.0
// @description This is a sample school API server.
// @host localhost:8081
// @BasePath /api
func main() {
	// Swagger documentation setup
	docs.SwaggerInfo.Title = "School API"
	docs.SwaggerInfo.Description = "This is a sample school API server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8081"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// Database connection
	dsn := "sqlserver://(local)?database=schooldb&trusted_connection=true&trustservercertificate=true&encrypt=false"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate the schema (only creates tables if they don't exist)
	err = db.AutoMigrate(&models.Class{}, &models.Student{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize repositories
	classRepo := repository.NewClassRepository(db)
	studentRepo := repository.NewStudentRepository(db)

	// Initialize services
	classService := service.NewClassService(classRepo)
	studentService := service.NewStudentService(studentRepo)

	// Initialize handlers
	classHandler := handler.NewClassHandler(classService)
	studentHandler := handler.NewStudentHandler(studentService)

	// Router setup
	router := mux.NewRouter()

	// Swagger documentation
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8081/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))

	// Class Routes
	router.HandleFunc("/api/classes", classHandler.CreateClass).Methods("POST")
	router.HandleFunc("/api/classes", classHandler.GetAllClasses).Methods("GET")
	router.HandleFunc("/api/classes/{id}", classHandler.GetClassByID).Methods("GET")
	router.HandleFunc("/api/classes/{id}", classHandler.UpdateClass).Methods("PUT")
	router.HandleFunc("/api/classes/{id}", classHandler.DeleteClass).Methods("DELETE")

	// Student Routes
	router.HandleFunc("/api/students", studentHandler.CreateStudent).Methods("POST")
	router.HandleFunc("/api/students", studentHandler.GetAllStudents).Methods("GET")
	router.HandleFunc("/api/students/{id}", studentHandler.GetStudentByID).Methods("GET")
	router.HandleFunc("/api/students/{id}", studentHandler.UpdateStudent).Methods("PUT")
	router.HandleFunc("/api/students/{id}", studentHandler.DeleteStudent).Methods("DELETE")

	// Start server in a goroutine
	go func() {
		log.Println("Server starting on port 8081...")
		log.Fatal(http.ListenAndServe(":8081", router))
	}()

	// Wait for server to start
	time.Sleep(1 * time.Second)

	// Open Swagger in default browser
	url := "http://localhost:8081/swagger/"
	var errOpen error
	switch runtime.GOOS {
	case "windows":
		errOpen = exec.Command("cmd", "/c", "start", url).Start()
	case "darwin":
		errOpen = exec.Command("open", url).Start()
	default:
		errOpen = exec.Command("xdg-open", url).Start()
	}
	if errOpen != nil {
		log.Printf("Failed to open browser: %v", errOpen)
	}

	// Keep the program running
	select {}
} 
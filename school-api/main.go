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
// @host localhost:8080
// @BasePath /api
func main() {
	// Swagger documentation setup
	docs.SwaggerInfo.Title = "School API"
	docs.SwaggerInfo.Description = "This is a sample school API server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// Database connection
	dsn := "sqlserver://(local)?database=schooldb&trusted_connection=true&trustservercertificate=true&encrypt=false"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Drop existing table
	err = db.Migrator().DropTable(&models.Class{})
	if err != nil {
		log.Println("Warning - Failed to drop table:", err)
	}

	// Auto Migrate the schema
	err = db.AutoMigrate(&models.Class{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize repository, service, and handler
	classRepo := repository.NewClassRepository(db)
	classService := service.NewClassService(classRepo)
	classHandler := handler.NewClassHandler(classService)

	// Router setup
	router := mux.NewRouter()

	// Swagger documentation
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))

	// Routes
	router.HandleFunc("/api/classes", classHandler.CreateClass).Methods("POST")
	router.HandleFunc("/api/classes", classHandler.GetAllClasses).Methods("GET")
	router.HandleFunc("/api/classes/{id}", classHandler.GetClassByID).Methods("GET")

	// Start server in a goroutine
	go func() {
		log.Println("Server starting on port 8080...")
		log.Fatal(http.ListenAndServe(":8080", router))
	}()

	// Wait for server to start
	time.Sleep(1 * time.Second)

	// Open Swagger in default browser
	url := "http://localhost:8080/swagger/"
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
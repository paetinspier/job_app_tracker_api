package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/paetinspier/gorest/application"
	"github.com/paetinspier/gorest/database"
	"github.com/paetinspier/gorest/middleware"
)

func InitializeModules(dbService *database.DatabaseService, router *http.ServeMux) {
	application.RegisterModule(dbService, router)
}

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	dbService, err := database.Init()
	if err != nil {
		fmt.Println(err)
	}

	if err := dbService.PingDatabase(); err != nil {
		fmt.Printf("Could not ping db: %v\n", err)
	}

	defer dbService.DB.Close()

	router := http.NewServeMux()

	InitializeModules(dbService, router)

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr: ":8080",
		Handler: stack(router),
	}
	color.HiBlue("Server running on port 8080")
	log.Fatal(server.ListenAndServe())
}

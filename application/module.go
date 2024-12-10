package application

import (
	"net/http"

	"github.com/paetinspier/gorest/database"
)

func RegisterModule(dbService *database.DatabaseService, router *http.ServeMux) {
	repo := NewRepository(dbService.DB)
	service := NewService(repo)
	handler := NewHandler(service, router)
	handler.InitRoutes()
}

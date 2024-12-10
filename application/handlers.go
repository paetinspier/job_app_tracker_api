package application

import (
	"net/http"

	"github.com/fatih/color"
)


type Handler struct {
	Service *Service
	Router *http.ServeMux
}

func NewHandler(service *Service, router *http.ServeMux) *Handler {
	return &Handler{Service: service, Router: router}
}

func (h *Handler) InitRoutes() {
	color.Green("Initializing Application Routes")
	h.Router.HandleFunc("/application", h.GetApplication)
}

func (h *Handler) GetApplication(w http.ResponseWriter, r *http.Request){
	color.Green("GET /application")
	application, err := h.Service.GetApplicationById(1)
	if err != nil {
		http.Error(w, "Application not found", http.StatusNotFound)
		return
	}

	w.Write([]byte(application.Title))
}

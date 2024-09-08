package handlers

import (
	"net/http"

	"github.com/OzodbekX/bookings/pkg/config"
	"github.com/OzodbekX/bookings/pkg/models"
	"github.com/OzodbekX/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remuteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remute_ip", remuteIp)


	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remuteIp := m.App.Session.GetString(r.Context(), "remute_ip")
	stringMap["remute_ip"] = remuteIp

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}

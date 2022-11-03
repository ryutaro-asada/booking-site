package handlers

import (
	"fmt"
	"net/http"

	"github.com/ryutaro-asada/go-practice/pkg/config"
	"github.com/ryutaro-asada/go-practice/pkg/models"
	"github.com/ryutaro-asada/go-practice/pkg/render"
)

// Repo is the pointer of repository used by the handlers
var Repo *Repository

// Repository is th repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		// App is point to config
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Sesstion.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// about is the about page hadler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	fmt.Println("inProduction")
	fmt.Println(m.App.InProduction)
	fmt.Println(r.RemoteAddr)
	// fmt.Println(r.Context().Value())
	remoteIP := m.App.Sesstion.GetString(r.Context(), "remote_ip")
	fmt.Println("aaaaaaaaa", remoteIP)
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

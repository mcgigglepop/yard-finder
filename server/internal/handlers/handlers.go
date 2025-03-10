package handlers

import (
	"log"
	"net/http"

	"github.com/mcgigglepop/yard-finder/server/internal/config"
	"github.com/mcgigglepop/yard-finder/server/internal/forms"
	"github.com/mcgigglepop/yard-finder/server/internal/models"
	"github.com/mcgigglepop/yard-finder/server/internal/render"
	"github.com/mcgigglepop/yard-finder/server/internal/repository"
	"github.com/mcgigglepop/yard-finder/server/internal/repository/dbrepo"
)

// repository used by the handlers
var Repo *Repository

// the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Index is the index/home page handler
func (m *Repository) Index(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "index.page.tmpl", &models.TemplateData{})
}

// GetCreateListing is create listing page handler for get requests
func (m *Repository) GetCreateListing(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "create-listing.page.tmpl", &models.TemplateData{})
}

func (m *Repository) PostCreateListing(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.RenewToken(r.Context())

	if err := r.ParseForm(); err != nil {
		log.Println("Error parsing form:", err)
		m.App.Session.Put(r.Context(), "error", "Error processing request")
		http.Redirect(w, r, "/create-listing", http.StatusSeeOther)
		return
	}

	// Validate form input
	form := forms.New(r.PostForm)
	// form.Required("projectName")

	if !form.Valid() {
		log.Println("Invalid form submission")
		render.Template(w, r, "create-listing.page.tmpl", &models.TemplateData{
			Form: form,
		})
		return
	}

	// // Construct project model
	// item := models.Project{
	// 	Name:      r.Form.Get("projectName"),
	// 	CreatedBy: 1,
	// 	Status:    "active",
	// }

	// // Insert into DB
	// _, err := m.DB.CreateProject(item)
	// if err != nil {
	// 	log.Println("Database insert error:", err)
	// 	m.App.Session.Put(r.Context(), "error", "Could not create project")
	// 	http.Redirect(w, r, "/projects", http.StatusSeeOther)
	// 	return
	// }

	m.App.Session.Put(r.Context(), "flash", "Listing created successfully")
	http.Redirect(w, r, "/create-listing", http.StatusSeeOther)
}

// GetConfirmListing is confirm listing page handler for get requests
func (m *Repository) GetConfirmListing(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "confirm-listing.page.tmpl", &models.TemplateData{})
}

// GetViewListing is view listing page handler for get requests
func (m *Repository) GetViewListing(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "view-listing.page.tmpl", &models.TemplateData{})
}
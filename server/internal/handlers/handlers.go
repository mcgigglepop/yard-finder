package handlers

import (
	"net/http"

	"github.com/mcgigglepop/yard-finder/server/internal/config"
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

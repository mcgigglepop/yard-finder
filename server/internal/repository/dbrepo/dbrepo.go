package dbrepo

import (
	"github.com/mcgigglepop/yard-finder/server/internal/config"
	"github.com/mcgigglepop/yard-finder/server/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
}

func NewPostgresRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
	}
}

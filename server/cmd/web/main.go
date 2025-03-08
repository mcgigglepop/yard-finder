package main

import (
	"encoding/gob"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/mcgigglepop/yard-finder/server/internal/config"
	"github.com/mcgigglepop/yard-finder/server/internal/handlers"
	"github.com/mcgigglepop/yard-finder/server/internal/helpers"
	"github.com/mcgigglepop/yard-finder/server/internal/models"
	"github.com/mcgigglepop/yard-finder/server/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	// Initialize application
	run()

	// Start the HTTP server
	log.Printf("Starting application on port %s", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	// Run the server
	log.Fatal(srv.ListenAndServe())
}

// run initializes the application
func run() {
	// Register models for session storage
	gob.Register(models.User{})
	gob.Register(map[string]int{})

	// Define command-line flags
	inProduction := flag.Bool("production", true, "Application is in production")
	useCache := flag.Bool("cache", true, "Use template cache")

	// Parse flags
	flag.Parse()

	// Set application config
	app.InProduction = *inProduction
	app.UseCache = *useCache

	// Set up logging
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.InfoLog = infoLog
	app.ErrorLog = errorLog

	// Configure session management
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	// Create template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc

	// Initialize handlers and utilities
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)
}

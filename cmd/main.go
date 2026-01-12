package main

import (
	"log"
	"net/http"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/raffidevaa/me-commerce/pkg/config"
	"github.com/raffidevaa/me-commerce/pkg/database"
	"gorm.io/gorm"
)

func handleCommand(args []string, db *gorm.DB) {
	for _, cmd := range args {
		switch cmd {
		case "migrate":
			log.Println("Running migration...")
			if err := database.AutoMigrate(db); err != nil {
				log.Fatal(err)
			}
			log.Println("Migration done")

		case "seed":
			log.Println("Running seeder...")
			if err := database.Seed(db); err != nil {
				log.Fatal(err)
			}
			log.Println("Seeder done")

		default:
			log.Fatalf("Unknown command: %s", cmd)
		}
	}
}

func main() {
	cfg := config.Load()

	db := database.NewPostgres(database.PostgresConfig{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
	})
	database.AutoMigrate(db)
	database.Seed(db)

	args := os.Args

	if len(args) > 1 {
		handleCommand(args[1:], db)
		return
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is healthy"))
	})
	http.ListenAndServe(":8080", r)
	log.Println("DB:", cfg.DBHost, cfg.DBPort)

	myFigure := figure.NewColorFigure("ME-COMMERCE", "", "green", true)
	myFigure.Print()

}

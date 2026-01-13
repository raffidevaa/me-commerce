package main

import (
	"log"
	"net/http"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/raffidevaa/me-commerce/internal/product"
	"github.com/raffidevaa/me-commerce/internal/user"
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

func loadConfiguration() (*gorm.DB, error) {
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

	return db, nil
}

func handleRoutes(r chi.Router, db *gorm.DB) {
	// user
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository, db)
	userController := user.NewUserController(userService)
	user.Routes(r, userController)

	// product
	productRepository := product.NewProductRepository(db)
	productService := product.NewProductService(productRepository, db)
	productController := product.NewProductController(productService)
	product.Routes(r, productController)
}

func main() {
	// load configuration
	db, err := loadConfiguration()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	args := os.Args

	if len(args) > 1 {
		handleCommand(args[1:], db)
		return
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Throttle(10))

	//handle routes
	handleRoutes(r, db)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is healthy"))
	})

	myFigure := figure.NewColorFigure("ME-COMMERCE", "", "green", true)
	myFigure.Print()

	log.Println("Server is speed-running on :8080")
	http.ListenAndServe(":8080", r)
}

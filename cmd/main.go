package main

import (
	"log"
	"net/http"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/raffidevaa/me-commerce/internal/cart"
	"github.com/raffidevaa/me-commerce/internal/product"
	"github.com/raffidevaa/me-commerce/internal/user"
	"github.com/raffidevaa/me-commerce/pkg/config"
	"github.com/raffidevaa/me-commerce/pkg/database"
	"github.com/raffidevaa/me-commerce/pkg/jwtauth"
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

func handleRoutes(r chi.Router, tokenAuth *jwtauth.JWTAuth, db *gorm.DB) {
	// user
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository, tokenAuth, db)
	userController := user.NewUserController(userService)
	user.Routes(r, userController)

	// product
	productRepository := product.NewProductRepository(db)
	productService := product.NewProductService(productRepository, db)
	productController := product.NewProductController(productService)
	product.Routes(r, productController)

	// cart
	cartRepository := cart.NewCartRepository(db)
	cartService := cart.NewCartService(cartRepository, db)
	cartController := cart.NewCartController(cartService)
	cart.Routes(r, cartController, tokenAuth)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
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

	// generate token jwt
	tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET_KEY")), nil)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Throttle(10))

	//handle routes
	handleRoutes(r, tokenAuth, db)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is healthy"))
	})

	myFigure := figure.NewColorFigure("ME-COMMERCE", "", "green", true)
	myFigure.Print()

	log.Println("Server is speed-running on 8080")
	http.ListenAndServe(":8080", r)
}

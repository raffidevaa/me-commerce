package cart

import (
	"github.com/go-chi/chi/v5"
	"github.com/raffidevaa/me-commerce/pkg/jwtauth"
)

type CartRoutes struct {
	CartController *CartController
}

func NewCartRoutes(cc *CartController) *CartRoutes {
	return &CartRoutes{
		CartController: cc,
	}
}

func Routes(r chi.Router, c *CartController, tokenAuth *jwtauth.JWTAuth) {
	r.Route("/carts", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Post("/items", c.AddItemToCart)
		r.Get("/", c.GetCartByUserID)
	})

}

package order

import (
	"github.com/go-chi/chi/v5"
	"github.com/raffidevaa/me-commerce/pkg/jwtauth"
)

type OrderRoutes struct {
	orderController *OrderController
}

func NewOrderRoutes(orderController *OrderController) *OrderRoutes {
	return &OrderRoutes{orderController: orderController}
}

func Routes(r chi.Router, c *OrderController, tokenAuth *jwtauth.JWTAuth) {
	r.Route("/orders", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Post("/checkout", c.CreateOrder)
		r.Get("/", c.GetOrdersByUserID)
	})

}

package user

import "github.com/go-chi/chi/v5"

type UserRoutes struct {
	UserController *UserController
}

func NewUserRoutes(ac *UserController) *UserRoutes {
	return &UserRoutes{
		UserController: ac,
	}
}

func Routes(r chi.Router, c *UserController) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", c.Register)
		r.Post("/login", c.Login)
	})

}

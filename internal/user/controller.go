package user

import (
	"encoding/json"
	"net/http"

	responsebuilder "github.com/raffidevaa/me-commerce/pkg/response-builder"
)

type UserController struct {
	userService *UserService
}

func NewUserController(us *UserService) *UserController {
	return &UserController{userService: us}
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responsebuilder.BadRequest(w, "invalid request body")
		return
	}

	user, err := c.userService.Register(r.Context(), req.ToUser())
	if err != nil {
		responsebuilder.BadRequest(w, err.Error())
		return
	}

	res := RegisterResponseFromUser(user)

	responsebuilder.Created(w, "user registered successfully", res)
}

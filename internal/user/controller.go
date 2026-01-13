package user

import (
	"encoding/json"
	"net/http"

	responseBuilder "github.com/raffidevaa/me-commerce/pkg/response-builder"
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
		responseBuilder.BadRequest(w, "invalid request body")
		return
	}

	user, err := c.userService.Register(r.Context(), req.ToUser())
	if err != nil {
		responseBuilder.BadRequest(w, err.Error())
		return
	}

	res := RegisterResponseFromUser(user)

	responseBuilder.Created(w, "user registered successfully", res)
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseBuilder.BadRequest(w, "invalid request body")
		return
	}

	loginRes, err := c.userService.Login(r.Context(), req)
	if err != nil {
		responseBuilder.Unauthorized(w, err.Error())
		return
	}

	responseBuilder.OK(w, "login successful", loginRes)
}

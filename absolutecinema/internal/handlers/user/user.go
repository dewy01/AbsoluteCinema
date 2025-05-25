package handler

import (
	"encoding/json"
	"net/http"

	"absolutecinema/internal/openapi/gen/usergen"
	user_service "absolutecinema/internal/service/user"
)

type UserHandler struct {
	Service user_service.Service
}

func NewUserHandler(svc user_service.Service) *UserHandler {
	return &UserHandler{Service: svc}
}

// POST /users/register
func (h *UserHandler) PostUsersRegister(w http.ResponseWriter, r *http.Request) {
	var input usergen.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	userOut, err := h.Service.Register(user_service.CreateUserInput{
		Name:            input.Name,
		Email:           input.Email,
		Password:        input.Password,
		ConfirmPassword: input.ConfirmPassword,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := usergen.UserOutput{
		Id:    &userOut.ID,
		Name:  &userOut.Name,
		Email: &userOut.Email,
		Role:  (*string)(&userOut.Role),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// POST /users/login
func (h *UserHandler) PostUsersLogin(w http.ResponseWriter, r *http.Request) {
	var input usergen.LoginUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := h.Service.Login(user_service.LoginInput{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	token := "dummy-token"

	resp := usergen.LoginResponse{
		Token: &token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// GET /users
func (h *UserHandler) GetUsersList(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

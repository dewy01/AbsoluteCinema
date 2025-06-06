package handler

import (
	"absolutecinema/internal/auth"
	"absolutecinema/internal/openapi/gen/usergen"
	user_service "absolutecinema/internal/service/user"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
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

	_, err := h.Service.Register(user_service.CreateUserInput{
		Name:            input.Name,
		Email:           input.Email,
		Password:        input.Password,
		ConfirmPassword: input.ConfirmPassword,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// POST /users/login
func (h *UserHandler) PostUsersLogin(w http.ResponseWriter, r *http.Request) {
	var input usergen.LoginUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, session, err := h.Service.Login(r.Context(), user_service.LoginInput{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	cookie := session.ToCookie(time.Now().Add(24 * time.Hour))
	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusNoContent)
}

// POST /users/logout
func (h *UserHandler) PostUsersLogout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(auth.CookieName)
	if err != nil {
		http.Error(w, "No active session", http.StatusUnauthorized)
		return
	}

	sessionID, err := uuid.Parse(cookie.Value)
	if err != nil {
		http.Error(w, "Invalid session", http.StatusBadRequest)
		return
	}

	if err := h.Service.Logout(r.Context(), sessionID); err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, auth.NewInvalidCookie())
	w.WriteHeader(http.StatusNoContent)
}

// GET /users/me
func (h *UserHandler) GetUsersMe(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(auth.CookieName)
	if err != nil {
		http.Error(w, "No active session", http.StatusUnauthorized)
		return
	}

	sessionID, err := uuid.Parse(cookie.Value)
	if err != nil {
		http.Error(w, "Invalid session", http.StatusBadRequest)
		return
	}

	user, err := h.Service.GetMe(r.Context(), sessionID)
	if err != nil {
		http.Error(w, "No user data found", http.StatusInternalServerError)
		return
	}

	resp := usergen.UserOutput{
		Id:    &user.ID,
		Name:  &user.Name,
		Email: &user.Email,
		Role:  (*string)(&user.Role),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// GET /users/{id}
func (h *UserHandler) GetUsersId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	userOut, err := h.Service.GetByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	resp := usergen.UserOutput{
		Id:    &userOut.ID,
		Name:  &userOut.Name,
		Email: &userOut.Email,
		Role:  (*string)(&userOut.Role),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// PUT /users/{id}
func (h *UserHandler) PutUsersId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	var input usergen.UpdateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	userOut, err := h.Service.Update(id, user_service.UpdateUserInput{
		Name:            *input.Name,
		Email:           *input.Email,
		Password:        *input.Password,
		ConfirmPassword: *input.ConfirmPassword,
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
	json.NewEncoder(w).Encode(resp)
}

// DELETE /users/{id}
func (h *UserHandler) DeleteUsersId(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	if err := h.Service.Delete(id); err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

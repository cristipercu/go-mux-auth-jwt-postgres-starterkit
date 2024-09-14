package user

import (
	"fmt"
	"net/http"

	"github.com/cristipercu/go-mux-auth-jwt-postgres-starterkit/cmd/config"
	"github.com/cristipercu/go-mux-auth-jwt-postgres-starterkit/service/auth"
	"github.com/cristipercu/go-mux-auth-jwt-postgres-starterkit/types"
	"github.com/cristipercu/go-mux-auth-jwt-postgres-starterkit/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
  store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
  return &Handler{
    store: store,
  }
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
  router.HandleFunc("/register", h.handleRegister).Methods(http.MethodPost)
}

func(h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
  var userPayload types.RegisterUserPayload

  if err := utils.ParseJSON(r, &userPayload); err != nil {
    utils.WriteError(w, http.StatusBadRequest, err)
  } 

  if err := utils.Validate.Struct(userPayload); err != nil {
    err := err.(validator.ValidationErrors)
    utils.WriteError(w, http.StatusBadRequest, err)
    return
  }

  _, err := h.store.GetUserByEmail(userPayload.Email)
  if err == nil {
    utils.WriteError(w, http.StatusConflict, fmt.Errorf("user %s aleready exists", userPayload.Email))
    return
  }

  hashedPassword, err := auth.HashPassword(userPayload.Password)
  if err != nil {
    utils.WriteError(w, http.StatusInternalServerError, err)
    return
  }

  err = h.store.CreateUser(types.User{
    Username: userPayload.Username,
    Email: userPayload.Email,
    Password: hashedPassword,
  })

  if err != nil {
    utils.WriteError(w, http.StatusBadRequest, err)
    return
  }

  utils.WriteJSON(w, http.StatusCreated, nil)
}

func(h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
  var loginPayload types.LoginUserPayload

  if err := utils.ParseJSON(r, &loginPayload); err != nil {
    utils.WriteError(w, http.StatusBadRequest, err)
    return
  }

  if err := utils.Validate.Struct(loginPayload); err != nil {
    errors := err.(validator.ValidationErrors)
    utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid payload: %v", errors))
    return
  }

  user, err := h.store.GetUserByEmail(loginPayload.Email)
  if err != nil {
    utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
    return
  }

  if !auth.ComparePassword(user.Password, loginPayload.Password) {
    utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
    return
  }
  
  secret := []byte(config.Envs.JWTSecret)
  token, err := auth.CreateJWT(secret, user.ID)

  if err != nil {
    utils.WriteError(w, http.StatusInternalServerError, err)
    return
  }

  utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
} 

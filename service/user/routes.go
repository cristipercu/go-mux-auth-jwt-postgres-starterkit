package user

import (
	"fmt"
	"net/http"

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

  err = h.store.CreateUser(userPayload)

  if err != nil {
    utils.WriteError(w, http.StatusBadRequest, err)
    return
  }

  utils.WriteJSON(w, http.StatusCreated, nil)
}

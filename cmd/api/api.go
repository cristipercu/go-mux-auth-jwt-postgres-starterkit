package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/cristipercu/go-mux-auth-jwt-postgres-starterkit/service/user"
	"github.com/gorilla/mux"
)

type APIserver struct {
  addr string
  db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIserver {
  return &APIserver{
    addr: addr,
    db: db,
  }
}


func (api *APIserver) Run() error {
  router := mux.NewRouter()
  subrouter := router.PathPrefix("/api/v1/").Subrouter()

  subrouter.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){
    w.WriteHeader(http.StatusOK)
  }).Methods(http.MethodGet)

  userStore := user.NewStore(api.db)
  userHandler := user.NewHandler(userStore)
  userHandler.RegisterRoutes(subrouter)

  log.Printf("Listening on %s", api.addr )

  return http.ListenAndServe(api.addr, router)
}

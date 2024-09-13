package user

import (
	"database/sql"

	"github.com/cristipercu/go-mux-auth-jwt-postgres-starterkit/types"
)

type Store struct {
  db *sql.DB
}

func NewStore (db *sql.DB) *Store {
  return &Store{
    db: db,
  }
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
  user := new(types.User)

  err := s.db.QueryRow(`SELECT id, username, email, password, created_on, modified_on
    FROM public.users WHERE email = $1`, email).Scan(
      &user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedOn, &user.ModifiedOn, 
    )

  if err != nil {
    return nil, err
  }

  return user, nil
}

func (s *Store) CreateUser(user types.RegisterUserPayload) error {
  _, err := s.db.Exec(`INSERT INTO public.users(username, email, password) VALUES ($1, $2, $3)`, user.Username, user.Email, user.Password)

  if err != nil {
    return err
  }
  return nil
}



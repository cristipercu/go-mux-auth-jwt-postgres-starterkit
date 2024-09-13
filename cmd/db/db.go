package db

import (
	"database/sql"
	"fmt"

  _ "github.com/lib/pq"
)


func NewPGStorage(user, dbname, password, sslmode string) (*sql.DB, error) {
  connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", user, dbname, password, sslmode)

  db, err := sql.Open("postgres", connStr)
  if err != nil {
    return nil, err
  }

  if err := db.Ping(); err != nil {
    return nil, err
  }

  return db, nil
} 

package main

import (
	"log"
	"os"

	"github.com/cristipercu/go-mux-auth-jwt-postgres-starterkit/cmd/config"
	"github.com/cristipercu/go-mux-auth-jwt-postgres-starterkit/cmd/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
  db, err := db.NewPGStorage(config.Envs.DbUser, config.Envs.DbName, config.Envs.DbPassword, config.Envs.Sslmode)

  if err != nil {
    log.Fatal(err)
  }

  driver, err := postgres.WithInstance(db, &postgres.Config{}) 
  if err != nil {
    log.Fatal(err)
  }

  m, err := migrate.NewWithDatabaseInstance("file://cmd/migrate/migrations", "postgres", driver)

  if err != nil {
    log.Fatal(err)
  }

  cmd := os.Args[(len(os.Args) - 1)]

  if cmd == "up" {
    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
      log.Fatal(err)
    }
  }

  if cmd == "down" {
    if err := m.Down(); err != nil && err != migrate.ErrNoChange {
      log.Fatal(err)
    }
  }
}

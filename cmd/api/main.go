package main

import (
	"log"

	"github.com/ThuraMinThein/social-golang/internal/db"
	"github.com/ThuraMinThein/social-golang/internal/env"
	"github.com/ThuraMinThein/social-golang/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr: 	   		env.GetString("DB_ADDR", "postgres://admin:password@localhost/social?sslmode=disable"),
			maxOpenConns:	env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: 	env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  	env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic("Failed to connect to database: ", err)
	}

	defer db.Close()
	log.Println("Connected to database successfully")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mont()

	log.Fatal(app.run(mux))
}
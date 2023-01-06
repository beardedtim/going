package data

import (
	"context"
	"fmt"

	"mck-p/modi/env"
	mcLog "mck-p/modi/log"

	"github.com/go-pg/pg/v10"
)

var db *pg.DB
var log mcLog.Logger = mcLog.CreateLogger("DATABASE")

// This function will make a connection to the database only once.
func Init() {
	var err error
	log.Info("Initializing and connecting to the database")

	if db != nil {
		log.Debug("Database already initalized and connected. No-op")
		return
	}

	log.Info("Creating new database connection")

	db = pg.Connect(&pg.Options{
		User: env.GetEnv("DB_USERNAME"),
		Addr: fmt.Sprintf(
			"%s:%s",
			env.GetEnv("DB_HOST"),
			env.GetEnv("DB_PORT"),
		),
		Password: env.GetEnv("DB_PASSWORD"),
		Database: env.GetEnv("DB_NAME"),
	})

	if err != nil {
		panic(err)
	}

	// this will be printed in the terminal, confirming the connection to the database
	log.Info("The database is connected and initalized.")
}

func Disconnect() error {
	log.Info("Disconnecting from Database")

	if db == nil {
		log.Debug("No Database instance to disconnect to.")

		return nil
	}

	return db.Close()
}

func IsHealthy() bool {
	// Not connected
	if db == nil {
		return false
	}

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		return false
	}

	// We are connected and we can ping
	return true
}

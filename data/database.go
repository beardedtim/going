package data

import (
	"context"

	"github.com/go-pg/pg/v10"
	mcLog "mck-p.com/log"
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
		User:     "username",
		Addr:     "0.0.0.0:9999",
		Password: "password",
		Database: "mckp",
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

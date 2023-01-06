package main

import (
	"os"
	"os/signal"

	database "mck-p/modi/data"
	ckp "mck-p/modi/http"
	mclog "mck-p/modi/log"
)

func main() {
	log := mclog.CreateLogger("MAIN")

	log.Info("Starting Database")

	database.Init()

	log.Info("Database Started. Starting server")

	server := ckp.CreateServer()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit

		if err := server.Stop(); err != nil {
			panic(err)
		}

		log.Info("Disconnected Server")

		if err := database.Disconnect(); err != nil {
			panic(err)
		}

		log.Info("Disconnected DB")

		os.Exit(0)
	}()

	server.Start()
}

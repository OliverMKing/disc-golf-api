package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	. "discgolfapi.com/m/server"
	"github.com/rs/zerolog/log"
)

// @title Disc Golf API
// @version 0.0.0

// @host localhost:8080
// @BasePath /
func main() {
	srv := GetServer()

	// start server
	go func() {
		log.Info().Msg("Starting http server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal().Msg(err.Error())
		}
	}()

	// handle graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c // block until signal received

	// wait for existing connections to finish
	wait := time.Second * 15
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	log.Info().Msg("Shutting down")
	srv.Shutdown(ctx)
	os.Exit(0)
}

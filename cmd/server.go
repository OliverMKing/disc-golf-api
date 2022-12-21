package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/olivermking/disc-golf-api/pkg/gen/openapi"
	"github.com/olivermking/disc-golf-api/pkg/service"
	"github.com/spf13/cobra"
)

func init() {
	f := serverCmd.Flags()
	f.IntVarP(&port, portFlag, "p", 0, "port that the server will run on (required)")
	serverCmd.MarkFlagRequired(portFlag)

	rootCmd.AddCommand(serverCmd)
}

var (
	serverCmd = &cobra.Command{
		Use:   "server [flags]",
		Short: "Starts the Disc Golf API server",
		Run:   RunServer,
	}

	portFlag = "port"
	port     int
)

const (
	timeout = 5 * time.Second
)

func RunServer(_ *cobra.Command, _ []string) {

	service := service.New()
	controller := openapi.NewDiscApiController(service)
	router := openapi.NewRouter(controller)

	addr := fmt.Sprintf(":%d", port)
	srv := &http.Server{Handler: router, Addr: addr, WriteTimeout: timeout, ReadTimeout: timeout}

	go func() {
		fmt.Printf("starting server on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal((err))
		}
	}()

	// waits for a sigint for a graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// shuts down the server after waiting for any final requests to finish
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down server")
}

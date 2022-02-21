package server

import (
	"flag"
	"net/http"
	"time"

	_ "discgolfapi.com/m/docs"
	"discgolfapi.com/m/server/middleware"

	"github.com/gorilla/mux"

	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	port = flag.String("port", "8080", "port")
)

func GetServer() *http.Server {
	// parse flags
	flag.Parse()

	router := mux.NewRouter()
	router.Use(middleware.Log)

	// api routes
	api := router.PathPrefix("/api").Subrouter()
	v1 := api.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/discs", GetDiscs).Methods("GET")
	v1.HandleFunc("/discs", PutDiscs).Methods("PUT")

	// doc routes
	router.PathPrefix("/").Handler((httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
	)))

	return &http.Server{
		Addr:         "0.0.0.0:" + *port,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
}

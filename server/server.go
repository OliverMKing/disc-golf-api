package server

import (
	"fmt"
	"net/http"
	"time"

	_ "discgolfapi.com/m/docs"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	httpSwagger "github.com/swaggo/http-swagger"
)

func GetServer() *http.Server {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	// api routes
	router.HandleFunc("/discs", GetDiscs).Methods("GET")
	router.HandleFunc("/discs", PutDiscs).Methods("PUT")

	// doc routes
	router.PathPrefix("/swagger/").Handler((httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
	)))

	return &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg(fmt.Sprintf("Received request to %s from %s", r.RequestURI, r.Host))
		next.ServeHTTP(w, r)
	})
}

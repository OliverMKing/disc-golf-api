package middleware

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg(fmt.Sprintf("Received request to %s %s from %s", r.Method, r.RequestURI, r.Host))
		next.ServeHTTP(w, r)
	})
}

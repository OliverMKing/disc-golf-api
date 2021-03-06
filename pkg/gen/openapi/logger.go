/*
 * Disc Golf API
 *
 * Open-source Disc golf api that follows OpenAPI specification. Source code can be found [here](https://github.com/OliverMKing/disc-golf-api).
 *
 * API version: 0.0.1
 * Contact: olivermerkleyking@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

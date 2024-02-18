// The syntax in this file requires Go >= 1.22
// more at https://www.willem.dev/articles/url-path-parameters-in-routes/#pattern-with-trailing-slash
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		response := fmt.Sprintf("returning users")
		w.Write([]byte(response))
	})

	// Fixing trailing slashes, to make life easier for API clients.
	// NOTE: be careful not all API clients do follow redirects per default
	mux.HandleFunc("GET /users/{$}", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/users", http.StatusSeeOther)
	})

	mux.HandleFunc("POST /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		userId := r.PathValue("id")
		response := fmt.Sprintf("Created or updated user %s", userId)
		w.Write([]byte(response))
	})

	log.Println("Listening to port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err) // most likely port is in use or no permissions because it is below 1024
	}
}

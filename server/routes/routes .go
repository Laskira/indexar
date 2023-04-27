package routes

import (
	"fmt"
	"net/http"

	"github.com/Laskira/indexar/server/request"
	"github.com/go-chi/chi/v5"
)

func StartRouting() {
	r := chi.NewRouter()

	r.Use(corsMiddleware)

	//Request to get the zincsearch data
	r.Get("/database", request.MappingRequest)

	fmt.Println("localhost running on port 3000")
	http.ListenAndServe(":3000", r)

}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if origin == "http://localhost:8081" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

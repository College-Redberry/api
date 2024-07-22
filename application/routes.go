package controller

import (
	"log"
	"net/http"

	"com.redberry.api/application/middlewares"
	"github.com/go-pkgz/routegroup"
)

func InitHttpRoutes() {
	router := routegroup.New(http.NewServeMux())
	router.Use(middlewares.Logging)

	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
		w.WriteHeader(http.StatusOK)
	})

	RoutesV1(router)

	log.Printf("Server listening on :8080")
	http.ListenAndServe(":8080", router)
}

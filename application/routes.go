package controller

import (
	"log"
	"net/http"

	"com.redberry.api/application/middlewares"
	_ "com.redberry.api/docs"
	"github.com/go-pkgz/routegroup"
	httpSwagger "github.com/swaggo/http-swagger"
)

func InitHttpRoutes() {
	router := routegroup.New(http.NewServeMux())
	router.Use(middlewares.Logging)

	router.HandleFunc("GET /swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
		w.WriteHeader(http.StatusOK)
	})

	RoutesV1(router)

	log.Printf("Server listening on :8080")
	http.ListenAndServe(":8080", router)
}

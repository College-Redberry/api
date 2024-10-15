package controller

import (
	"com.redberry.api/application/controller/auth"
	"com.redberry.api/application/controller/projects"
	"com.redberry.api/application/controller/users"
	"com.redberry.api/application/middlewares"
	"github.com/go-pkgz/routegroup"
)

func RoutesV1(router *routegroup.Bundle) {
	authController := auth.New()
	usersController := users.New()
	projectsController := projects.New()

	// Endpoints
	apiRouter := router.Mount("/api/v1")

	authRouter := apiRouter.Mount("/auth")
	authRouter.HandleFunc("POST /login", authController.Login)

	// From here, needs the auth token
	apiRouter.Use(middlewares.Auth)

	usersRouter := apiRouter.Mount("/users")
	usersRouter.HandleFunc("POST ", usersController.Register)
	projectsRouter := apiRouter.Mount("/projects")
	projectsRouter.HandleFunc("POST ", projectsController.Create)
	projectsRouter.HandleFunc("GET /{projectID}", projectsController.GetByID)
	projectsRouter.HandleFunc("PUT /{projectID}", projectsController.Update)
	projectsRouter.HandleFunc("DELETE /{projectID}", projectsController.Delete)


	// From here, needs to be admin
	apiRouter.Use(middlewares.Permission)
}

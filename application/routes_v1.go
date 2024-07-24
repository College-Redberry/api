package controller

import (
	"com.redberry.api/application/controller/auth"
	"com.redberry.api/application/controller/users"
	"github.com/go-pkgz/routegroup"
)

func RoutesV1(router *routegroup.Bundle) {
	apiRouter := router.Mount("/api/v1")

	authController := auth.New()
	authRouter := apiRouter.Mount("/auth")
	authRouter.HandleFunc("POST /login", authController.Login)

	usersController := users.New()
	usersRouter := apiRouter.Mount("/users")
	usersRouter.HandleFunc("POST ", usersController.Register)
}

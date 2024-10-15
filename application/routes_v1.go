package controller

import (
	"com.redberry.api/application/controller/auth"
	"com.redberry.api/application/controller/boards"
	"com.redberry.api/application/controller/cards"
	"com.redberry.api/application/controller/messages"
	"com.redberry.api/application/controller/priorities"
	"com.redberry.api/application/controller/projects"
	"com.redberry.api/application/controller/statuses"
	"com.redberry.api/application/controller/users"
	"com.redberry.api/application/middlewares"
	"github.com/go-pkgz/routegroup"
)

func RoutesV1(router *routegroup.Bundle) {
	authController := auth.New()
	usersController := users.New()
	projectsController := projects.New()
	boardsController := boards.New()
	statusesController := statuses.New()
	prioritiesController := priorities.New()
	cardsController := cards.New()
	messagesController := messages.New()

	// Endpoints
	apiRouter := router.Mount("/api/v1")

	authRouter := apiRouter.Mount("/auth")
	authRouter.HandleFunc("POST /login", authController.Login)

	// From here, needs the auth token
	apiRouter.Use(middlewares.Auth)

	usersRouter := apiRouter.Mount("/users")
	usersRouter.HandleFunc("POST ", usersController.Register)
	usersRouter.HandleFunc("GET /", usersController.GetByEmail)
	usersRouter.HandleFunc("UPDATE /{userID}", usersController.Update)
	usersRouter.HandleFunc("DELETE /{userID}", usersController.Delete)

	projectsRouter := apiRouter.Mount("/projects")
	projectsRouter.HandleFunc("POST ", projectsController.Create)
	projectsRouter.HandleFunc("GET /{projectID}", projectsController.GetByID)
	projectsRouter.HandleFunc("PUT /{projectID}", projectsController.Update)
	projectsRouter.HandleFunc("DELETE /{projectID}", projectsController.Delete)

	boardsRouter := apiRouter.Mount("/boards")
	boardsRouter.HandleFunc("POST ", boardsController.Create)
	boardsRouter.HandleFunc("GET /{boardID}", boardsController.GetByID)
	boardsRouter.HandleFunc("PUT /{boardID}", boardsController.Update)
	boardsRouter.HandleFunc("DELETE /{boardID}", boardsController.Delete)

	statusesRouter := apiRouter.Mount("/statuses")
	statusesRouter.HandleFunc("POST ", statusesController.Create)
	statusesRouter.HandleFunc("GET /{statusID}", statusesController.GetByID)
	statusesRouter.HandleFunc("PUT /{statusID}", statusesController.Update)
	statusesRouter.HandleFunc("DELETE /{statusID}", statusesController.Delete)

	prioritiesRouter := apiRouter.Mount("/priorities")
	prioritiesRouter.HandleFunc("POST ", prioritiesController.Create)
	prioritiesRouter.HandleFunc("GET /{priorityID}", prioritiesController.GetByID)
	prioritiesRouter.HandleFunc("PUT /{priorityID}", prioritiesController.Update)
	prioritiesRouter.HandleFunc("DELETE /{priorityID}", prioritiesController.Delete)

	cardsRouter := apiRouter.Mount("/cards")
	cardsRouter.HandleFunc("POST ", cardsController.Create)
	cardsRouter.HandleFunc("GET /{cardID}", cardsController.GetByID)
	cardsRouter.HandleFunc("PUT /{cardID}", cardsController.Update)
	cardsRouter.HandleFunc("DELETE /{cardID}", cardsController.Delete)

	messagesRouter := apiRouter.Mount("/messages")
	messagesRouter.HandleFunc("POST ", messagesController.Create)
	messagesRouter.HandleFunc("GET /{messageID}", messagesController.GetByID)
	messagesRouter.HandleFunc("PUT /{messageID}", messagesController.Update)
	messagesRouter.HandleFunc("DELETE /{messageID}", messagesController.Delete)


	// From here, needs to be admin
	apiRouter.Use(middlewares.Permission)
}

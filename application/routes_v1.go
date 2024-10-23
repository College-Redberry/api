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
	// apiRouter.Use(middlewares.Auth)

	usersRouter := apiRouter.Mount("/users")
	usersRouter.HandleFunc("POST ", usersController.Register)
	usersRouter.HandleFunc("GET /", usersController.GetByEmail)
	usersRouter.HandleFunc("PUT /{user_id}", usersController.Update)
	usersRouter.HandleFunc("DELETE /{user_id}", usersController.Delete)

	projectsRouter := apiRouter.Mount("/projects")
	projectsRouter.HandleFunc("POST ", projectsController.Create)
	projectsRouter.HandleFunc("GET /{project_id}", projectsController.GetByID)
	projectsRouter.HandleFunc("PUT /{project_id}", projectsController.Update)
	projectsRouter.HandleFunc("DELETE /{project_id}", projectsController.Delete)

	boardsRouter := apiRouter.Mount("/boards")
	boardsRouter.HandleFunc("POST ", boardsController.Create)
	boardsRouter.HandleFunc("GET /{board_id}", boardsController.GetByID)
	boardsRouter.HandleFunc("PUT /{board_id}", boardsController.Update)
	boardsRouter.HandleFunc("DELETE /{board_id}", boardsController.Delete)

	statusesRouter := apiRouter.Mount("/statuses")
	statusesRouter.HandleFunc("POST ", statusesController.Create)
	statusesRouter.HandleFunc("GET /{status_id}", statusesController.GetByID)
	statusesRouter.HandleFunc("PUT /{status_id}", statusesController.Update)
	statusesRouter.HandleFunc("DELETE /{status_id}", statusesController.Delete)

	prioritiesRouter := apiRouter.Mount("/priorities")
	prioritiesRouter.HandleFunc("POST ", prioritiesController.Create)
	prioritiesRouter.HandleFunc("GET /{priority_id}", prioritiesController.GetByID)
	prioritiesRouter.HandleFunc("PUT /{priority_id}", prioritiesController.Update)
	prioritiesRouter.HandleFunc("DELETE /{priority_id}", prioritiesController.Delete)

	cardsRouter := apiRouter.Mount("/cards")
	cardsRouter.HandleFunc("POST ", cardsController.Create)
	cardsRouter.HandleFunc("GET /{card_id}", cardsController.GetByID)
	cardsRouter.HandleFunc("PUT /{card_id}", cardsController.Update)
	cardsRouter.HandleFunc("DELETE /{card_id}", cardsController.Delete)

	messagesRouter := apiRouter.Mount("/messages")
	messagesRouter.HandleFunc("POST ", messagesController.Create)
	messagesRouter.HandleFunc("GET /{message_id}", messagesController.GetByID)
	messagesRouter.HandleFunc("PUT /{message_id}", messagesController.Update)
	messagesRouter.HandleFunc("DELETE /{message_id}", messagesController.Delete)

	// From here, needs to be admin
	apiRouter.Use(middlewares.Permission)
}

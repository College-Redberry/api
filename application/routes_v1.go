package controller

import (
	"com.redberry.api/application/controller/auth"
	"com.redberry.api/application/controller/messages" // Importando o controlador de mensagens
	"com.redberry.api/application/controller/users"
	"com.redberry.api/application/middlewares"
	"github.com/go-pkgz/routegroup"
)

func RoutesV1(router *routegroup.Bundle) {
	authController := auth.New()
	usersController := users.New()
	messagesController := messages.NewMessagesController() // Instanciando o controlador de mensagens

	// Endpoints
	apiRouter := router.Mount("/api/v1")

	authRouter := apiRouter.Mount("/auth")
	authRouter.HandleFunc("POST /login", authController.Login)

	// A partir daqui, é necessário o token de autenticação
	apiRouter.Use(middlewares.Auth)

	usersRouter := apiRouter.Mount("/users")
	usersRouter.HandleFunc("POST /register", usersController.Register)

	// A partir daqui, é necessário ser administrador
	apiRouter.Use(middlewares.Permission)

	// Rotas para mensagens
	messagesRouter := apiRouter.Mount("/messages")
	messagesRouter.HandleFunc("POST /", messagesController.Create)    // Criar mensagem
	messagesRouter.HandleFunc("GET /", messagesController.GetAll)     // Listar todas as mensagens
	messagesRouter.HandleFunc("GET /get", messagesController.GetByID) // Recuperar mensagem por ID
	messagesRouter.HandleFunc("PUT /", messagesController.Update)     // Atualizar mensagem
	messagesRouter.HandleFunc("DELETE /", messagesController.Delete)  // Deletar mensagem
}

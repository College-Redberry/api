package main

import controller "com.redberry.api/application"

// @title Redberry API
// @version 1.0
// @description Project management system.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
    controller.InitHttpRoutes()
}

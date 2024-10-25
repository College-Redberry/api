package main

import controller "com.redberry.api/application"

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
    controller.InitHttpRoutes()
}

package main

import "amar_dokan/cmd"

// @title Amar Dokan API
// @version 1.0
// @description API documentation for Amar Dokan.
// @host localhost:8000
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	cmd.Serve()
}

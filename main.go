package main

import (
	"github.com/stevo1403/go-by-example/apps"
)

func init() {
	Load()
}

// @title UnGo API
// @version 1.0
// @description This is a simple API for a blogging platform
// @BasePath /api/v1
// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description "Bearer token for API authorization"
// @description Type "Bearer" followed by a space and JWT token.
// @schemes http
// @produce json
// @consumes json
// externalDocs.description OpenAPI
func main() {
	apps.LoadViews()
}

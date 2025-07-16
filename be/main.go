package main

import (
	"github.com/nchikkam/context-findr-be/router"
	"github.com/nchikkam/context-findr-be/utils"
)

// @Author:   nchikkam
// @title           Simple File Upload Go OpenAPI
// @version         1.0
// @description     A simple file Upload CRUD API using Gin.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host            localhost:8080
// @BasePath        /
func main() {
	server := router.SetUpServer()
	server.Run(utils.Port)
}

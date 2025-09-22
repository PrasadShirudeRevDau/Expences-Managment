package main

import (
	config "ExpencesManagment/Config"
	routes "ExpencesManagment/Routes"
	swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.DatabaseConnection()

	r:=routes.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
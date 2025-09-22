package routes

import (
	controller "ExpencesManagment/Controller"
	"ExpencesManagment/docs"
	

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r:=gin.Default()

	docs.SwaggerInfo.BasePath="/"
	r.POST("/expenses",controller.CreateExpenses)

	return r
}
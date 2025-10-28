package routes

import (
	controller "ExpencesManagment/Controller"
	middlewares "ExpencesManagment/Middlewares"
	"ExpencesManagment/docs"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	r.POST("/auth/register", controller.RegisterUser)
	r.POST("/auth/Login", controller.LoginUser)

	authorized := r.Group("/api", middlewares.AuthMiddleware())
	{
		authorized.POST("/expenses", controller.CreateExpenses)
		authorized.DELETE("/expenses/:id", controller.DeleteExpences)
		authorized.GET("/expenses/all", controller.GetAllExpences)
		authorized.GET("/expenses/:id", controller.GetExpenseById)
		authorized.GET("/expenses/filter", controller.GetExpenseByFilter)
		authorized.GET("/expenses/category/all", controller.GetCategory)
		authorized.PATCH("/expenses/:id", controller.UpdateExpences)

	}

	return r
}

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
					//Auth 
	r.POST("/auth/register",middlewares.RegisterUser)
	r.POST("/auth/Login", middlewares.LoginUser)

	authorized := r.Group("/", middlewares.AuthMiddleware())
	{
		authorized.POST("/expenses", controller.CreateExpenses)
		authorized.DELETE("/expenses/:id", controller.DeleteExpences)
		authorized.GET("/expenses/all", controller.GetAllExpences)
		authorized.GET("/expenses/:id", controller.GetExpenseById)
		authorized.GET("/expenses/filter", controller.GetExpenseByFilter)
		authorized.GET("/expenses/category/all", controller.GetCategory)
		authorized.PATCH("/expenses/:id", controller.UpdateExpences)
		authorized.GET("/analytics/category", controller.TotalByCategory)
		authorized.GET("/analytics/range", controller.TotalByDateRange)
		authorized.GET("/analytics/summary", controller.SummaryOfExpenses)
		authorized.GET("analytics/permonth", controller.SummaryPerMonth)
	}

	return r
}

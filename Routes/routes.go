package routes

import (
	controller "ExpencesManagment/Controller"
	"ExpencesManagment/docs"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	r.POST("/expenses", controller.CreateExpenses)
	r.DELETE("/expenses/:id", controller.DeleteExpences)
	r.GET("/expenses/all", controller.GetAllExpences)
	r.GET("/expenses/:id", controller.GetExpenseById)
	r.GET("/expenses/date",controller.GetExpenseByDate)
	r.GET("/expenses/category",controller.GetExpensesByCategory)
	r.GET("/expenses/category/all",controller.GetCategory)
	r.PATCH("/expenses/:id", controller.UpdateExpences)

	return r
}

package controller

import (
	models "ExpencesManagment/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)
//GetAllExpenses godoc
//@Summary show all categories 
//@Description get all categories
//@Tags categories
//@Accept json
//@Produce json
//@Security BearerAuth
//@Router /api/expenses/category/all [get]
func GetCategory(c *gin.Context) {
	categories:=models.GetAllCategories()
	c.JSON(http.StatusOK,gin.H{"category": categories})
}
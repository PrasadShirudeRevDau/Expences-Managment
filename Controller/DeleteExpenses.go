package controller

import (
	config "ExpencesManagment/Config"
	models "ExpencesManagment/Models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
//DeleteExpenses godoc
//@Summary delete expense
//@Description delete expense by id
//@Tags expenses
//@Accept json
//@Produce json
//@Security BearerAuth
//@Param id path string true "expense id"
//@Router /api/expenses/{id} [delete]
func DeleteExpences(c *gin.Context) {
	id:=c.Param("id")
	var expense models.Expense

	ExpenseID, err :=strconv.Atoi(id)
	if err !=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid expense id"})
		return 
	}

	if err:=config.DatabaseConnection().Where("id = ?", ExpenseID).First(&expense).Error;err !=nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}
	if err :=config.DatabaseConnection().Delete(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to delete expenses"})
		return 
	}
	c.JSON(http.StatusOK,"Expense delete succesfully")

}
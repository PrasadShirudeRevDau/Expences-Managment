package controller

import (
	config "ExpencesManagment/Config"
	models "ExpencesManagment/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateAmount struct {
	Amount float64 `json:"amount"`
}
//UpdateExpenses godoc
//@Summary update expense amount
//@Description update amount by id
//@Tags expenses
//@Accept json
//@Produce json
//@Param id path string true "Expense ID"
//@Param body body UpdateAmount true "New amount"
//@Router /expenses/{id} [patch]
func UpdateExpences(c *gin.Context) {

	db := config.DatabaseConnection()
	id := c.Param("id")
	var input UpdateAmount

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Model(&models.Expense{}).Where("id = ?", id).Update("amount", input.Amount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Expense updated successfully"})
}

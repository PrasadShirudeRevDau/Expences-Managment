package controller

import (
	config "ExpencesManagment/Config"
	models "ExpencesManagment/Models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateInput struct {
	Amount   float64
	Category string
	Date     string
	Note     string
}
//CreateExpenses godoc
//@Summary Create expense
//@Description Create an new expenses
//@Tags expenses
//@Accept json
//@Produce json
//@Param body body CreateInput true "New expense"
//@Router /expenses [post]
func CreateExpenses(c *gin.Context) {
	db := config.DatabaseConnection()
	var input CreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parseDate, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Invalid date formate. use YYYY-MM-DD"})
		return
	}

	ValidCateories := []models.ExpenseCategory{
		models.Clothing,
		models.Education,
		models.Entertainment,
		models.Food_Groceries,
		models.Gifts_Donation,
		models.Health_Fitness,
		models.Housing,
		models.Insurance,
		models.Miscellaneous,
		models.Personal_Care,
		models.Transportation,
		models.Travel_Vacation,
		models.Utilities,
	}

	isValid := false
	for _, cat:= range ValidCateories {
		if models.ExpenseCategory(input.Category) ==cat {
			isValid= true
			break	
		}
	}

	if !isValid {
		c.JSON(http.StatusBadRequest,gin.H{"error": "Invalid category"})
		return
	}

	expense := models.Expense{
		Amount:   input.Amount,
		Category: models.ExpenseCategory(input.Category),
		Date:     parseDate,
		Note:     input.Note,
	}

	if err := db.Create(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to create expence"})
		return
	}

	c.JSON(http.StatusOK, expense)
}

package controller

import (
	"ExpencesManagment/Models"
	"ExpencesManagment/Config"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

type CreateInput struct {
	Amount float64
	Category string
	Date string
	Note string
}


func CreateExpenses(c *gin.Context) {
	db:=config.DatabaseConnection()
	var input CreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	parseDate , err :=time.Parse("2006-01-02",input.Date)
	if err !=nil {
		c.JSON(http.StatusBadGateway,gin.H{"error":"Invalid date formate. use YYYY-MM-DD"})
		return
	}

	expense:= models.Expense {
		Amount: input.Amount,
		Category: input.Category,
		Date: parseDate,
		Note: input.Note,		
	}

	if err:= db.Create(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Faild to create expence"})
		return
	}

	c.JSON(http.StatusOK, expense)
}

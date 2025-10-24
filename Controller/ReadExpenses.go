package controller

import (
	config "ExpencesManagment/Config"
	models "ExpencesManagment/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var expenses []models.Expense
var db = config.DatabaseConnection()
//GetAllExpenses godoc
//@Summary read all expense
//@Description get all expenses
//@Tags expenses
//@Accept json
//@Produce json
//@Router /expenses/all [get]
func GetAllExpences(c *gin.Context) {

	if err := db.Find(&expenses); err != nil {
		c.JSON(http.StatusOK, expenses)
		return
	}
}
//GetExpenseById godoc
//@Summary get single expense
//@Description get single expense by id
//@Tags expenses
//@Accept json
//@Produce json
//@Param id path string true "Expense ID"
//@Router /expenses/{id} [get]
func GetExpenseById(c *gin.Context) {
	id := c.Param("id")
	var expense models.Expense

	if err := db.First(&expense, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}

	c.JSON(http.StatusOK, expense)
	fmt.Println("Expense ID:", expense.Id)
}

//@Summary get expense by date
//@Description get expenses by date
//@Tags expenses
//@Accept json
//@Produce json
//@Param from query string true "Start date (YYYY-MM-DD)"
//@Param to query string true "End date (YYYY-MM-DD)"
//@Router /expenses/date [get]
func GetExpenseByDate(c *gin.Context) {
	From:=c.Query("from")
	To:=c.Query("to")

	if From == "" || To ==""{
		c.JSON(http.StatusBadRequest,gin.H{"error":"From and To dates are required"})
		return 
	}
	var expense []models.Expense

	if err:=db.Where("date between ? and ?",From,To).Find(&expense).Error; err !=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	if len(expense)==0 {
		c.JSON(http.StatusNotFound,gin.H{"error":"No Expense founded"})
		return
	}

	c.JSON(http.StatusOK, expense)
}

//@Summary get expense by category
//@Description user can get expense by category
//@Tags expenses
//@Accept json
//@Produce json
//@Param category query string true "category"
//@Router /expenses/category [get]
func GetExpensesByCategory (c *gin.Context) {
	category:=c.Query("category")
	
	if category==""{
		c.JSON(http.StatusNotFound,gin.H{"error":"Category is empty"})
		return
	}
	var expense []models.Expense

	isValid:=false
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

	for _,ValidCategory := range ValidCateories{
		if ValidCategory == models.ExpenseCategory(category){
			isValid=true
			break
		}
	}

	if !isValid {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid category"})
		return 
	}

	if err := db.Where("category = ?",category).Find(&expense).Error; err !=nil {
		c.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, expense)
}

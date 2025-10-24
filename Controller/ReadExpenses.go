package controller

import (
	config "ExpencesManagment/Config"
	models "ExpencesManagment/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = config.DatabaseConnection()

// GetAllExpenses godoc
// @Summary read all expense
// @Description get all expenses
// @Tags expenses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Router /api/expenses/all [get]
func GetAllExpences(c *gin.Context) {
	userID := c.GetUint("user_id")
	var expense []models.Expense
	result := db.Where("user_id =?", userID).Find(&expense)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, expense)
}

// GetExpenseById godoc
// @Summary get single expense
// @Description get single expense by id
// @Tags expenses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Expense ID"
// @Router /api/expenses/{id} [get]
func GetExpenseById(c *gin.Context) {
	id := c.Param("id")
	var expense models.Expense

	if err := db.First(&expense, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}

	c.JSON(http.StatusOK, expense)
}

// @Summary get expense
// @Description get expenses by date and category
// @Tags expenses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string false "id"
// @Param from query string false "Start date (YYYY-MM-DD)"
// @Param to query string false "End date (YYYY-MM-DD)"
// @Param category query string false "category"
// @Router /api/expenses/filter [get]
func GetExpenseByFilter(c *gin.Context) {
	id := c.Query("id")
	From := c.Query("from")
	To := c.Query("to")
	category := c.Query("category")

	var expenses []models.Expense
	query := db.Model(&models.Expense{})

	if id != "" {
		query = query.Where("id = ?", id)
	}

	if From != "" && To != "" {
		query = query.Where("date between ? and ?", From, To)
	} else if From != "" || To != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both dates are required"})
		return
	}

	if category != "" {
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

		for _, ValidCategory := range ValidCateories {
			if ValidCategory == models.ExpenseCategory(category) {
				isValid = true
				break
			}
		}

		if !isValid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
			return
		}

		query = query.Where("category = ?", category)
	}

	if id == "" && From == "" && To == "" && category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error":"enter atlest one valid value"})
		return
	}

	if err := query.Find(&expenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(expenses) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Expense founded"})
		return
	}

	c.JSON(http.StatusOK, expenses)

}
